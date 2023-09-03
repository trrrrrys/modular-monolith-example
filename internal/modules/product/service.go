package product

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/product/v1"
)

func NewproductService() product.ProductServiceServer {
	return &productService{}
}

type productService struct{}

func (productService) GetItem(ctx context.Context, in *product.GetItemRequest) (*product.GetItemResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("ProductServce/service.GetItem").End()
	time.Sleep(20 * time.Millisecond)
	return &product.GetItemResponse{
		ItemId:     in.ItemId,
		Price:      1000,
		StockCount: 10,
	}, nil
}

func (productService) ShipItem(ctx context.Context, r *product.ShipItemRequest) (*product.ShipItemResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("ProductServce/service.ShihpItem").End()
	time.Sleep(20 * time.Millisecond)
	return &product.ShipItemResponse{}, nil
}
