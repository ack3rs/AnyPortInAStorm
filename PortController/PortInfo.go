package PortController

import (
	"encoding/json"

	l "github.com/acky666/ackyLog"
)

type PortInfo struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type Ports = map[string]PortInfo

func UnMarshalWholeMap(rawJSON []byte) (Ports, error) {

	Out := map[string]PortInfo{}

	err := json.Unmarshal(rawJSON, &Out)
	if err != nil {
		l.ERROR("JSON Decode Failed: %v", err)
	}

	for k, v := range Out {
		l.INFO("PortName: %s  %+v", k, v)
	}
	return Out, nil
}

func UnMarshalPortInfo(rawJSON []byte) (PortInfo, error) {

	Out := PortInfo{}

	err := json.Unmarshal(rawJSON, &Out)
	if err != nil {
		l.ERROR("JSON Decode Failed: %v", err)
	}

	return Out, nil
}
