package rajaongkir

type Response struct {
	StatusCode  float64
	Description string
	Result      any
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

type ProvinceResponse struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}

type CityResponse struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type SubdistrictResponse struct {
	SubdistrictID   string `json:"subdistrict_id"`
	ProvinceID      string `json:"province_id"`
	Province        string `json:"province"`
	CityID          string `json:"city_id"`
	City            string `json:"city"`
	Type            string `json:"type"`
	SubdistrictName string `json:"subdistrict_name"`
}
