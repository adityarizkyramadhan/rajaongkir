package rajaongkir_test

import (
	"testing"

	"github.com/adityarizkyramadhan/rajaongkir"
)

func TestEnvironment(t *testing.T) {
	testCases := []struct {
		Name     string
		Type     rajaongkir.Type
		Expected string
	}{
		{
			Name:     "Starter Account",
			Type:     rajaongkir.Starter,
			Expected: "https://api.rajaongkir.com/starter",
		},
		{
			Name:     "Basic Account",
			Type:     rajaongkir.Basic,
			Expected: "https://api.rajaongkir.com/basic",
		},
		{
			Name:     "Pro Account",
			Type:     rajaongkir.Pro,
			Expected: "https://pro.rajaongkir.com/api",
		},
		{
			Name:     "Undefined",
			Type:     1000,
			Expected: "UNDEFINED",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := tc.Type.String()
			if result != tc.Expected {
				t.Errorf("Expected: %v, Got: %v", tc.Expected, result)
			}
		})
	}
}
