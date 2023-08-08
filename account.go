package rajaongkir

import "net/http"

type Account struct {
	client      *http.Client
	apiKey      string
	typeAccount Type
}

func New(apiKey string) *Account {
	return &Account{
		client: new(http.Client),
		apiKey: apiKey,
	}
}
