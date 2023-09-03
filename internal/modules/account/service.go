package account

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/account/v1"
)

func NewAccountService() account.AccountServiceServer {
	return &accountService{}
}

type accountService struct{}

func (a *accountService) GetUserInfo(ctx context.Context, in *account.GetUserInfoRequest) (*account.GetUserInfoResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("AccountServce/service.GetUserInfo").End()
	time.Sleep(20 * time.Millisecond)
	return &account.GetUserInfoResponse{
		Id:   in.Id,
		Name: "name_" + in.Id,
	}, nil
}

func (a *accountService) GetShippingAddress(ctx context.Context, r *account.GetShippingAddressRequest) (*account.GetShippingAddressResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("AccountServce/service.GetShippingAddress").End()
	time.Sleep(20 * time.Millisecond)
	return &account.GetShippingAddressResponse{
		UserId:     r.UserId,
		PostalCode: "100-0005",
		Address:    "東京都千代田区丸の内１丁目",
	}, nil
}
