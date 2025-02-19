package service

import (
	"context"
	"github.com/cocopanda-dev/gorder-v2/common/metrics"
	"github.com/cocopanda-dev/gorder-v2/order/adapters"
	"github.com/cocopanda-dev/gorder-v2/order/app"
	"github.com/cocopanda-dev/gorder-v2/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricsClient),
		},
	}
}
