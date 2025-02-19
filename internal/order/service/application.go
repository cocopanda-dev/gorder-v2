package service

import (
	"context"
	"github.com/cocopanda-dev/gorder-v2/order/app"
)

func NewApplication(ctx context.Context) app.Application {
	//orderRepo := adapters.NewMemoryOrderRepository()
	return app.Application{}
}
