package commands

import (
	"context"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/config"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/internal/order/aggregate"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/es"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/logger"

	"github.com/pkg/errors"
)

type CreateOrderCommandHandler interface {
	Handle(ctx context.Context, command *aggregate.CreateOrderCommand) error
}

type createOrderHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewCreateOrderHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *createOrderHandler {
	return &createOrderHandler{log: log, cfg: cfg, es: es}
}

func (c *createOrderHandler) Handle(ctx context.Context, command *aggregate.CreateOrderCommand) error {
	err := c.es.Exists(ctx, command.AggregateID)
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}
	order := aggregate.NewOrderAggregateWithID(command.AggregateID)
	if err := order.HandleCommand(command); err != nil {
		return err
	}
	return c.es.Save(ctx, order)
}
