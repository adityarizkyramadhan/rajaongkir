package rajaongkir

import (
	"encoding/json"
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
	jsonData, err := json.Marshal(resp.Result)
	if err != nil {
		return nil, err
	}
	var provinces []ProvinceResponse
	if err := json.Unmarshal(jsonData, &provinces); err != nil {
		return nil, err
	}
	return provinces, nil
}
