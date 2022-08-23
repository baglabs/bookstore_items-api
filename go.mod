module github.com/baglabs/bookstore_items-api

go 1.18

require (
	github.com/baglabs/bookstore_oauth-go v0.0.0-20220823033529-3418d3149f98
	github.com/baglabs/bookstore_utils-go v0.0.0-20220823042319-dbba8f22f2fd
)

require github.com/gorilla/mux v1.8.0

replace github.com/baglabs/bookstore_oauth-go => ../bookstore_oauth-go
