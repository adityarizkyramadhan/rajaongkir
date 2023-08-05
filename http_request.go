package rajaongkir

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (a *Account) httpRequest(method, link, auth string, bodyReq io.Reader) (*Response, error) {
	req, err := http.NewRequest(method, link, bodyReq)
	if err != nil {
		return nil, err
	}
	req.Header.Set("key", auth)
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	status := data["rajaongkir"].(map[string]any)["status"].(map[string]any)["code"].(float64)
	description := data["rajaongkir"].(map[string]any)["status"].(map[string]any)["description"].(string)
	if status != 200 {
		return nil, errors.New(description)
	}

	result := data["rajaongkir"].(map[string]any)["results"]
	response := &Response{
		StatusCode:  status,
		Description: description,
		Result:      result,
	}
	return response, nil

}
