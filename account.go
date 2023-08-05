package rajaongkir

import "net/http"

type Account struct {
	client *http.Client
	Type   Type
}

func New() *Account {
	return &Account{
		client: new(http.Client),
	}
}
