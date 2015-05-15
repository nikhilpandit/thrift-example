package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/nikhilpandit/thrift-example/gen-go/hello"
	"github.com/nikhilpandit/thrift-example/service/db"
	"github.com/nikhilpandit/thrift-example/service/handler"
)

// Flags
var (
	host = flag.String("host", "localhost", "host to run the thrift server")
	port = flag.Int("port", 9090, "port to run the thrift server")
)

func main() {
	flag.Parse()
	// setup the database
	mongoURL, found := syscall.Getenv("MONGO_URL")
	if !found {
		log.Fatal("Need to set MONGO_URL")
	}
	log.Printf("Connecting to MongoDB at %s", mongoURL)
	database, err := db.NewMongoDB(mongoURL)
	if err != nil {
		log.Fatalf("Error connecting to mongo at %v. %v", mongoURL, err.Error())
	}
	defer database.Close()
	log.Printf("Connected to MongoDB at %s", mongoURL)

	// thrift magic begins here:
	listenAddr := fmt.Sprintf("%s:%d", *host, *port)
	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		log.Fatal("Error starting server socket at %s: %s", listenAddr, err)
	}
	defer serverTransport.Close()

	h := handler.NewHelloHandler(database)
	processor := hello.NewHelloProcessor(h)
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)

	log.Printf("Starting hello server listening on %s", listenAddr)
	if err = server.Serve(); err != nil {
		log.Fatalf("Error calling serve on hello server: %s", err.Error())
	}
}
