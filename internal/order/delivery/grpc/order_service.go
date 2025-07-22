package grpc

import (
	"context"

	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/logger"
	orderService "github.com/MateusLeviDev/Go-EventSourcing-CQRS/proto/order"
)

type orderGrpcService struct {
	log logger.Logger
}

func NewOrderGrpcService(log logger.Logger) *orderGrpcService {
	return &orderGrpcService{log: log}
}

func (o *orderGrpcService) CreateOrder(ctx context.Context, r *orderService.CreateOrderReq) (*orderService.CreateOrderRes, error) {
	//TODO
	panic("Implement")
}

func (o *orderGrpcService) PayOrder(ctx context.Context, r *orderService.PayOrderReq) (*orderService.PayOrderRes, error) {
	//TODO
	panic("Implement")
}

func (o *orderGrpcService) SubmitOrder(ctx context.Context, req *orderService.SubmitOrderReq) (*orderService.SubmitOrderRes, error) {
	//TODO
	panic("Implement")
}

func (o *orderGrpcService) GetOrderByID(ctx context.Context, req *orderService.GetOrderByIDReq) (*orderService.GetOrderByIDRes, error) {
	//TODO
	panic("Implement")
}

func (o *orderGrpcService) UpdateOrder(ctx context.Context, req *orderService.UpdateOrderReq) (*orderService.UpdateOrderRes, error) {
	//TODO
	panic("Implement")
}
