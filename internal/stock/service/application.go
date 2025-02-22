package service

import (
	"context"
	"github.com/cocopanda-dev/gorder-v2/common/metrics"
	"github.com/cocopanda-dev/gorder-v2/stock/adapters"
	"github.com/cocopanda-dev/gorder-v2/stock/app"
	"github.com/cocopanda-dev/gorder-v2/stock/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(_ context.Context) app.Application {
	orderRepo := adapters.NewMemoryStockRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TodoMetrics{}

	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			CheckIfItemsInStock: query.NewCheckIfItemsInStockHandler(orderRepo, logger, metricsClient),
			GetItems:            query.NewGetItemsHandler(orderRepo, logger, metricsClient),
		},
	}
}
