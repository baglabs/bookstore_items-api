module github.com/baglabs/bookstore_items-api

go 1.18

require github.com/baglabs/bookstore_utils-go v0.0.0-20220824134108-b54c1b9ad07e

require (
	github.com/baglabs/bookstore_oauth-go v0.0.0-20220823080415-a45b7f2f8d39
	github.com/gorilla/mux v1.8.0
)

require github.com/elastic/go-elasticsearch/v7 v7.17.1

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.22.0 // indirect
)

replace github.com/baglabs/bookstore_oauth-go => ../bookstore_oauth-go

replace github.com/baglabs/bookstore_utils-go => ../bookstore_utils-go
