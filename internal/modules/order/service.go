package order

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/account/v1"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/order/v1"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/payment/v1"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/product/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewOrderService(
	accountService account.AccountServiceClient,
	paymentService payment.PaymentServiceClient,
	productService product.ProductServiceClient,
) order.OrderServiceServer {
	return &orderService{
		accountService: accountService,
		paymentService: paymentService,
		productService: productService,
	}
}

type orderService struct {
	accountService account.AccountServiceClient
	paymentService payment.PaymentServiceClient
	productService product.ProductServiceClient
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func (o orderService) OrderItem(ctx context.Context, in *order.OrderItemRequest) (*order.OrderItemResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("OrderServce/service.OrderItem").End()
	orderID := uuid.New()

	item, _ := o.productService.GetItem(ctx, &product.GetItemRequest{
		ItemId: in.ItemId,
	})

	if item.StockCount < 1 {
		return nil, status.Errorf(codes.FailedPrecondition, "在庫不足")
	}
	// 発送状態
	_, _ = o.productService.ShipItem(ctx, &product.ShipItemRequest{})

	// 住所取得
	userAddress, _ := o.accountService.GetShippingAddress(ctx, &account.GetShippingAddressRequest{
		UserId: in.UserId,
	})

	// 決済実行
	_, _ = o.paymentService.PaymentItem(ctx, &payment.PaymentItemRequest{
		UserId: in.UserId,
		Price:  item.Price,
	})

	// 注文情報保存
	_ = userAddress

	return &order.OrderItemResponse{
		OrderId: orderID.String(),
		ItemId:  in.ItemId,
	}, nil
}
