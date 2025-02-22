package command

import (
	"context"

	"github.com/cocopanda-dev/gorder-v2/common/decorator"
	domain "github.com/cocopanda-dev/gorder-v2/order/domain/order"
	"github.com/sirupsen/logrus"
)

type UpdateOrder struct {
	order    *domain.Order
	updateFn func(context.Context, *domain.Order) (*domain.Order, error)
}

type UpdateOrderHandler decorator.CommandHandler[UpdateOrder, interface{}]

type updateOrderHandler struct {
	orderRepo domain.Repository
	//stockGRPC
}

func NewUpdateOrderHandler(
	orderRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) UpdateOrderHandler {
	if orderRepo == nil {
		panic("nil orderRepo")
	}
	return decorator.ApplyCommandDecorators[UpdateOrder, interface{}](
		updateOrderHandler{
			orderRepo: orderRepo,
		},
		logger,
		metricsClient,
	)
}

func (c updateOrderHandler) Handle(ctx context.Context, cmd UpdateOrder) (interface{}, error) {
	if cmd.updateFn == nil {
		logrus.Warnf("updateOrderHandler got nil UpdateFn, order=%#v", cmd.order)
		cmd.updateFn = func(ctx context.Context, order *domain.Order) (*domain.Order, error) { return order, nil }
	}
	if err := c.orderRepo.Update(ctx, cmd.order, cmd.updateFn); err != nil {
		return nil, err
	}
	return nil, nil
}
