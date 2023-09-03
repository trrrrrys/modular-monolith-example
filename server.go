package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	accountModule "github.com/trrrrrys/modular-monolith-example/internal/modules/account"
	orderModule "github.com/trrrrrys/modular-monolith-example/internal/modules/order"
	paymentModule "github.com/trrrrrys/modular-monolith-example/internal/modules/payment"
	productModule "github.com/trrrrrys/modular-monolith-example/internal/modules/product"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/account/v1"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/order/v1"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/payment/v1"
	"github.com/trrrrrys/modular-monolith-example/internal/proto/product/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func dummyInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	seg := newrelic.FromContext(ctx).StartSegment(
		fmt.Sprintf("%s/Interceptor", info.FullMethod),
	)
	time.Sleep(20 * time.Millisecond)
	seg.End()
	return handler(ctx, req)
}

func runServer() {

	nrenabled := os.Getenv("NEWRELIC_LICENSE_KEY") != ""
	nrapp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("modular-monolith-example"),
		newrelic.ConfigLicense(os.Getenv("NEWRELIC_LICENSE_KEY")),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(cfg *newrelic.Config) {
			cfg.HostDisplayName = "modular-monolith-example"
		},
		newrelic.ConfigEnabled(nrenabled),
	)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			nrgrpc.UnaryServerInterceptor(nrapp),
			dummyInterceptor,
		),
	)
	conn, err := grpc.Dial(
		"127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(nrgrpc.StreamClientInterceptor),
	)
	if err != nil {
		panic(err)
	}
	accountClient := account.NewAccountServiceClient(conn)
	paymentClient := payment.NewPaymentServiceClient(conn)
	productClient := product.NewProductServiceClient(conn)
	accountModule := accountModule.NewAccountService()
	productModule := productModule.NewproductService()
	paymentModule := paymentModule.NewpaymentService()

	orderModule := orderModule.NewOrderService(
		accountClient,
		paymentClient,
		productClient,
	)
	account.RegisterAccountServiceServer(s, accountModule)
	product.RegisterProductServiceServer(s, productModule)
	payment.RegisterPaymentServiceServer(s, paymentModule)
	order.RegisterOrderServiceServer(s, orderModule)

	reflection.Register(s)
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(fmt.Errorf("network I/O error: %w", err))
	}
	log.Println("start server")
	log.Fatal(s.Serve(lis))
}
