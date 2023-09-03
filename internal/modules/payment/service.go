package payment

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/payment/v1"
)

func NewpaymentService() payment.PaymentServiceServer {
	return &paymentService{}
}

type paymentService struct{}

func (paymentService) PaymentItem(ctx context.Context, in *payment.PaymentItemRequest) (*payment.PaymentItemResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("PaymentServce/service.PaymentItem").End()
	time.Sleep(20 * time.Millisecond)
	return &payment.PaymentItemResponse{
		Result: "ok",
	}, nil
}
