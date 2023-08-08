package rajaongkir

import (
	"fmt"
	"strings"
)

type pro struct {
	account *Account
}

func (a *Account) Pro() *pro {
	a.typeAccount = Pro
	return &pro{account: a}
}

func (p *pro) GetProvince() ([]ProvinceResponse, error) {
	resp, err := p.account.httpRequest("GET", p.account.typeAccount.String()+"/province", p.account.apiKey, nil)
	if err != nil {
		return nil, err
	}
	var provinces []ProvinceResponse
	if err := interfaceToStruct(resp.Result, &provinces); err != nil {
		return nil, err
	}
	return provinces, nil
}

func (p *pro) GetCityByIDProvince(idProvince string) ([]CityResponse, error) {
	resp, err := p.account.httpRequest("GET", p.account.typeAccount.String()+"/city?province="+idProvince, p.account.apiKey, nil)
	if err != nil {
		return nil, err
	}
	var cities []CityResponse
	if err := interfaceToStruct(resp.Result, &cities); err != nil {
		return nil, err
	}
	return cities, nil
}

func (p *pro) GetSubdistrictByIDCity(idCity string) ([]SubdistrictResponse, error) {
	resp, err := p.account.httpRequest("GET", p.account.typeAccount.String()+"/subdistrict?city="+idCity, p.account.apiKey, nil)
	if err != nil {
		return nil, err
	}
	var subdistricts []SubdistrictResponse
	if err := interfaceToStruct(resp.Result, &subdistricts); err != nil {
		return nil, err
	}
	return subdistricts, nil
}

func (p *pro) GetCost(input CostInputPro) ([]CostResponse, error) {
	payload := strings.NewReader(fmt.Sprintf("origin=%s&destination=%s&weight=%f&courier=%s&originType=%s&destinationType=%s", input.Origin, input.Destination, input.Weight, input.Courier, input.OriginType, input.DestinationType))
	resp, err := p.account.httpRequest("POST", p.account.typeAccount.String()+"/cost", p.account.apiKey, payload)
	if err != nil {
		return nil, err
	}
	var costs []CostResponse
	if err := interfaceToStruct(resp.Result, &costs); err != nil {
		return nil, err
	}
	return costs, nil
}
