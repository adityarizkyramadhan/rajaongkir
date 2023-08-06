package rajaongkir

import (
	"fmt"
	"strings"
)

type starter struct {
	account *Account
}

func (a *Account) Starter() *starter {
	a.typeAccount = Starter
	return &starter{account: a}
}

type ProvinceResponse struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}

func (s *starter) GetProvince() ([]ProvinceResponse, error) {
	resp, err := s.account.httpRequest("GET", s.account.typeAccount.String()+"/province", s.account.apiKey, nil)
	if err != nil {
		return nil, err
	}
	var provinces []ProvinceResponse
	if err := interfaceToStruct(resp.Result, &provinces); err != nil {
		return nil, err
	}
	return provinces, nil
}

type CityResponse struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

func (s *starter) GetCityByIDProvince(idProvince string) ([]CityResponse, error) {
	resp, err := s.account.httpRequest("GET", s.account.typeAccount.String()+"/city?province="+idProvince, s.account.apiKey, nil)
	if err != nil {
		return nil, err
	}
	var cities []CityResponse
	if err := interfaceToStruct(resp.Result, &cities); err != nil {
		return nil, err
	}
	return cities, nil
}

type CostResponse struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Costs []struct {
		Service     string `json:"service"`
		Description string `json:"description"`
		Cost        []struct {
			Value int    `json:"value"`
			Etd   string `json:"etd"`
			Note  string `json:"note"`
		} `json:"cost"`
	} `json:"costs"`
}

type CostInput struct {
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Weight      float64 `json:"weight"`
	Courier     string  `json:"courier"`
}

func (s *starter) GetCost(input CostInput) ([]CostResponse, error) {
	fmt.Println(input)
	payload := strings.NewReader(fmt.Sprintf("origin=%s&destination=%s&weight=%f&courier=%s", input.Origin, input.Destination, input.Weight, input.Courier))
	resp, err := s.account.httpRequest("POST", s.account.typeAccount.String()+"/cost", s.account.apiKey, payload)
	if err != nil {
		return nil, err
	}
	var costs []CostResponse
	if err := interfaceToStruct(resp.Result, &costs); err != nil {
		return nil, err
	}
	return costs, nil
}
