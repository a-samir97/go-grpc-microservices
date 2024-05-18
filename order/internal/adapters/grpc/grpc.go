package grpc

import (
	"context"

	"Github.com/a-samir97/microservices/order/internal/application/core/domain"
	"github.com/huseyinbabal/microservices-proto/golang/order"
)


func(a Adapter) Create(ctx context.Context, reqest *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems [] domain.OrderItem

	for _, orderItem := range reqest.OrderItems{
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice: orderItem.UnitPrice,
			Quantity: orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(reqest.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}