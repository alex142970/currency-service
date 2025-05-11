package main

import (
	"currency_service/currency/internal/config"
	"currency_service/currency/internal/db"
	"currency_service/currency/internal/handler"
	"currency_service/currency/internal/repository"
	"currency_service/currency/internal/service"
	pb "currency_service/pkg/generated/currency"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"os"
)

var (
	port       = flag.Int("port", 50051, "The server port")
	configPath = flag.String("config", "config.yaml", "Path to config file")
)

func main() {
	flag.Parse()

	logger := slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelDebug)

	conf, err := config.LoadConfig(*configPath)

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connection, err := db.NewDatabaseConnection(&conf.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := repository.NewCurrencyPostgresRepository(connection)

	currencyServer := handler.NewCurrencyServer(service.NewCurrencyService(repo, conf, logger), logger)

	err = startGRPCServer(currencyServer)
	if err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}

func startGRPCServer(server *handler.CurrencyServer) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCurrencyServer(s, server)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
