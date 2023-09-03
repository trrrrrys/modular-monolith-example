# modular-monolith-example
## Usage

```
# run server
$ go run .

# gRPC Request
$ grpcurl -plaintext \
    -d '{ "item_id": "item_1", "user_id": "user_1" }' \
    localhost:50051 order.v1.OrderService.OrderItem | jq
```

## Structure
```
.
├── buf.gen.yaml
├── buf.work.yaml
├── internal
│   ├── modules
│   └── proto
├── main.go
├── proto
│   ├── account
│   ├── buf.yaml
│   ├── order
│   ├── payment
│   └── product
└── server.go
```
