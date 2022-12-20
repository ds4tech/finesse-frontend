module github.com/ds4tech/covantis-sre

go 1.17

replace github.com/ds4tech/covantis-sre/pkg/common => ./pkg/common

require (
	github.com/ds4tech/covantis-sre/pkg/common v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	go.opentelemetry.io/otel v1.11.2 // indirect
	go.opentelemetry.io/otel/sdk v1.11.2 // indirect
	go.opentelemetry.io/otel/trace v1.11.2 // indirect
)
