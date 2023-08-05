package rajaongkir

type Type int

const (
	Starter Type = iota
	Basic   Type = iota
	Pro     Type = iota
)

func (t Type) String() string {
	switch t {
	case Starter:
		return "https://api.rajaongkir.com/starter"
	case Basic:
		return "https://api.rajaongkir.com/basic"
	case Pro:
		return "https://pro.rajaongkir.com/api"
	default:
		return "UNDEFINED"
	}
}
