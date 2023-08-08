package rajaongkir

type CostInput struct {
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Weight      float64 `json:"weight"`
	Courier     string  `json:"courier"`
}
