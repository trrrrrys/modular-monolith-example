FROM golang:1.21 AS builder
WORKDIR /go/src/github.com/trrrrrys/modular-monolith-example
COPY go.* .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ./bin/app .

FROM gcr.io/distroless/static-debian12
COPY --from=builder /go/src/github.com/trrrrrys/modular-monolith-example/bin/app /usr/local/bin/app
CMD ["app"]
