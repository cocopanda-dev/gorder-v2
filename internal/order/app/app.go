package app

import (
	"github.com/cocopanda-dev/gorder-v2/order/app/command"
	"github.com/cocopanda-dev/gorder-v2/order/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateOrder command.CreateOrderHandler
	UpdateOrder command.UpdateOrderHandler
}

type Queries struct {
	GetCustomerOrder query.GetCustomerOrderHandler
}
