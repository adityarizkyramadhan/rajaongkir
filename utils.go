package rajaongkir

import "encoding/json"

func interfaceToStruct(data, target any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, target); err != nil {
		return err
	}
	return nil
}
