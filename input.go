package rajaongkir

type CostInputStarter struct {
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Weight      float64 `json:"weight"`
	Courier     string  `json:"courier"`
}

type CostInputPro struct {
	Origin          string  `json:"origin"`
	OriginType      string  `json:"originType"`
	Destination     string  `json:"destination"`
	DestinationType string  `json:"destinationType"`
	Weight          float64 `json:"weight"`
	Courier         string  `json:"courier"`
}
