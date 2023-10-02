module github.com/ds4tech/webserver

go 1.19

replace github.com/ds4tech/covantis-sre/pkg/common => ./pkg/common

require (
	github.com/ds4tech/covantis-sre/pkg/common v0.0.0-20230102140711-722a093ce77b
	go.opentelemetry.io/otel v1.19.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/sdk v1.19.0
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	go.opentelemetry.io/otel/trace v1.19.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
)
