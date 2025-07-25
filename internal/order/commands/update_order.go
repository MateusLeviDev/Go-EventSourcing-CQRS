package commands

import (
	"context"

	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/config"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/internal/order/aggregate"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/es"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/logger"
)

type UpdateOrderCommandHandler interface {
	Handle(ctx context.Context, command *aggregate.OrderUpdatedCommand) error
}

type updateOrderCmdHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewUpdateOrderCmdHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *updateOrderCmdHandler {
	return &updateOrderCmdHandler{log: log, cfg: cfg, es: es}
}

func (c *updateOrderCmdHandler) Handle(ctx context.Context, command *aggregate.OrderUpdatedCommand) error {
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
