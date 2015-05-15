package main

import (
	"fmt"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/nikhilpandit/thrift-example/gen-go/hello"
)

func main() {
	// setup the thrift client
	var transport thrift.TTransport
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		log.Fatal("Error opening socket: ", err)
	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()

	if err = transport.Open(); err != nil {
		log.Fatal(err)
	}

	client := hello.NewHelloClientFactory(transport, protocolFactory)

	// make calls to the thrift service
	greeting, err := client.Hello("np")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greeting)
}
