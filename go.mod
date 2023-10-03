module github.com/ds4tech/webserver

go 1.21.1

replace github.com/ds4tech/finesse-frontend/pkg/common => ./pkg/common

require github.com/ds4tech/finesse-frontend/pkg/common v0.0.0-20231003084215-ebf6e70db24f

require github.com/gorilla/mux v1.8.0 // indirect
