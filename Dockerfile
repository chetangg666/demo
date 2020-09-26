FROM golang:latest

RUN apt-get update && apt-get install -y default-mysql-client
ENV mySQLIPAddress 127.0.0.1
ENV mySQLIPPort 3306


RUN echo $mySQLIPAddress
RUN echo $mySQLIPPort
RUN echo $mySQLPassword
#RUN apk --no-cache add ca-certificates
RUN go get github.com/go-sql-driver/mysql

RUN mkdir /app

ADD . /app

WORKDIR /app/app

# Fetch dependencies.
RUN go mod download

# we run go build to compile the binary
# executable of our Go program
RUN go build -o main ./cmd

# our newly created binary executable
CMD ["/app/app/main"]