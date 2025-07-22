package service

import (
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/config"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/internal/order/commands"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/internal/order/queries"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/es"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/logger"
)

type OrderService struct {
	Commands *commands.OrderCommands
	Queries  *queries.OrderQueries
}

func NewOrderService(
	log logger.Logger,
	cfg *config.Config,
	es es.AggregateStore,
	//mongoRepo repository.OrderRepository,
) *OrderService {
	createOrderHandler := commands.NewCreateOrderHandler(log, cfg, es)
	orderPaidHandler := commands.NewOrderPaidHandler(log, cfg, es)
	submitOrderHandler := commands.NewSubmitOrderHandler(log, cfg, es)
	updateOrderCmdHandler := commands.NewUpdateOrderCmdHandler(log, cfg, es)
	getOrderByIDHandler := queries.NewGetOrderByIDHandler(log, cfg, es)
	orderCommands := commands.NewOrderCommands(createOrderHandler, orderPaidHandler, submitOrderHandler, updateOrderCmdHandler)
	orderQueries := queries.NewOrderQueries(getOrderByIDHandler)
	return &OrderService{Commands: orderCommands, Queries: orderQueries}
}
