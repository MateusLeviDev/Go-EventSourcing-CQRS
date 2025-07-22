package commands

import (
	"context"

	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/config"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/internal/order/aggregate"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/es"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/logger"
)

type OrderPaidCommandHandler interface {
	Handle(ctx context.Context, command *aggregate.OrderPaidCommand) error
}

type orderPaidHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewOrderPaidHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *orderPaidHandler {
	return &orderPaidHandler{log: log, cfg: cfg, es: es}
}

func (c *orderPaidHandler) Handle(ctx context.Context, command *aggregate.OrderPaidCommand) error {
	err := c.es.Exists(ctx, command.AggregateID)
	if err != nil {
		return err
	}
	order := aggregate.NewOrderAggregateWithID(command.AggregateID)
	if err := c.es.Load(ctx, order.GetID(), order); err != nil {
		return err
	}
	if err := order.HandleCommand(command); err != nil {
		return err
	}
	return c.es.Save(ctx, order)
}
