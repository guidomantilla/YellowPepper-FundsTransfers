module YellowPepper-FundsTransfers

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gin-gonic/gin v1.7.3
	go.uber.org/zap v1.19.0

	go.opentelemetry.io/otel v1.0.0-RC1
    go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC1
    go.opentelemetry.io/otel/sdk v1.0.0-RC1
    go.opentelemetry.io/otel/trace v1.0.0-RC1

    github.com/nats-io/nats.go v1.12.0
)
