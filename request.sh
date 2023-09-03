
payload=`cat << EOF
{
  "item_id": "item_1",
  "user_id": "user_1"
}
` 
grpcurl -plaintext \
  -d "`echo $payload | jq`" \
  localhost:50051 order.v1.OrderService.OrderItem
