module github.com/ds4tech/covantis-sre/cmd

go 1.17

replace github.com/ds4tech/covantis-sre/pkg/common => ../pkg/common

require github.com/ds4tech/covantis-sre/pkg/common v0.0.0-00010101000000-000000000000

require github.com/gorilla/mux v1.8.0 // indirect
