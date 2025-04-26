package main

import (
	application "currency_service/currency/internal/app"
	"currency_service/currency/internal/handler"
	"currency_service/currency/internal/service"
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	app := application.NewApp()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		panic(err)
	}

	server := handler.NewServer(service.NewCurrencyService(app))

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}

	log.Printf("server listening at %v", lis.Addr())
}
