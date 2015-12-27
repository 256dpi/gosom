package gosom

import (
	"encoding/json"
	"io/ioutil"
)

func Store(som *SOM, file string) error {
	data, err := Encode(som)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, data, 0644)
}

func Load(file string) (*SOM, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return Decode(data)
}

func Encode(som *SOM) ([]byte, error) {
	return json.MarshalIndent(som, "", "  ")
}

func Decode(data []byte) (*SOM, error) {
	som := NewSOM(0, 0)
	err := json.Unmarshal(data, som)
	if err != nil {
		return nil, err
	}

	return som, nil
}
