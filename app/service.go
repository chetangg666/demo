package app

import (
	"context"
	"database/sql"
	"app/rabbitmqq"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// Service provides some "capabilities" to our application
type Service interface {
	Create(ctx context.Context, req addRequest) (int, error)
	Get(ctx context.Context) (string, error)
	Delete(ctx context.Context, date int) (int, error)
}

var (
	Id        int
	LastName  string
	FirstName string
)

type opService struct {
	db *sql.DB
	ch *amqp.Channel
	q  amqp.Queue
}

// NewService makes a new Service.
func NewService(db *sql.DB, ch *amqp.Channel, q amqp.Queue) Service {
	return &opService{db: db, ch: ch, q: q}
}

// create will insert data in db
func (db *opService) Create(ctx context.Context, req addRequest) (int, error) {
	stmt, err := db.db.Prepare("INSERT INTO Persons(LastName, FirstName) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(req.FirstName, req.LastName)
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	req.ID = int(lastId)
	req.MsgType = "Added"
	jsonData, err := json.Marshal(req)
	err = db.ch.Publish(
		"",        // exchange
		db.q.Name, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonData,
		})
	log.Printf(" [x] Sent %s", req)
	rabbitmqq.FailOnError(err, "Failed to publish a message")
	return req.ID, nil
}

// Get will return today's date
func (opService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02/01/2006"), nil
}

// Delete will  delete row from db
func (db *opService) Delete(ctx context.Context, ID int) (int, error) {
	var req addRequest
	///////////// Read Operation ///////////////////
	rowsQuery, err := db.db.Prepare("SELECT * FROM Persons where PersonID = ?")
	if err != nil {
		return ID, err
	}

	defer rowsQuery.Close()
	// Parametrize
	rows, err := rowsQuery.Query(ID)
	for rows.Next() {
		err := rows.Scan(&req.ID, &req.LastName, &req.FirstName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(req.ID, req.LastName, req.FirstName)
	}
	_, err = db.db.Exec("DELETE FROM Persons where PersonID = ?", ID) // OK
	if err != nil {
		log.Fatal(err)
	}
	if req.ID == 0 {
		return ID, errors.New("Not exist in db to delete")
	}
	req.MsgType = "Deleted"
	jsonData, err := json.Marshal(req)
	err = db.ch.Publish(
		"",        // exchange
		db.q.Name, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonData,
		})
	log.Printf(" [x] Sent %s", req)
	rabbitmqq.FailOnError(err, "Failed to publish a message")
	return ID, nil
}
