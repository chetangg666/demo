package main

import (
	"context"
	"demo"
	"demo/db"
	"demo/metrics"
	"demo/rabbitmqq"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		httpAddr = flag.String("http", ":9090", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	db := db.DataBase()
	defer db.Close()
	ch, q, conn := rabbitmqq.Config()
	defer conn.Close()
	// our  service
	srv := demo.NewService(db, ch, q)
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(metrics.RequestsToIndex)
	// mapping endpoints
	endpoints := demo.Endpoints{
		GetEndpoint:    demo.MakeGetEndpoint(srv),
		CreateEndpoint: demo.MakeCreateEndpoint(srv),
		DeleteEndpoint: demo.MakeDeleteEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("demo is listening on port:", *httpAddr)
		handler := demo.NewHTTPServer(ctx, endpoints, r)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
