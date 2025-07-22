package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/config"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/internal/order/service"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/es/store"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/eventstroredb"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/interceptors"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/logger"
	"github.com/go-playground/validator"
)

type server struct {
	cfg *config.Config
	log logger.Logger
	db  *esdb.Client
	im  interceptors.InterceptorManager
	os  *service.OrderService
	v   *validator.Validate
}

func NewServer(cfg *config.Config, log logger.Logger) *server {
	return &server{cfg: cfg, log: log, v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)

	db, err := eventstroredb.NewEventStoreDB(s.cfg.EventStoreConfig)
	if err != nil {
		return err
	}
	aggregateStore := store.NewAggregatorStore(s.log, db)
	s.os = service.NewOrderService(s.log, s.cfg, aggregateStore)

	closeGrpcServer, grpcServer, err := s.newOrderGrpcServer()

	if err != nil {
		return err
	}

	defer closeGrpcServer() // nolint: errcheck

	<-ctx.Done()
	grpcServer.GracefulStop()
	s.log.Info("Order server exited properly")
	return nil
}
