package rajaongkir

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
