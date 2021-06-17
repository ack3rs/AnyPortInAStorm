package PortController

import (
	"testing"

	l "github.com/acky666/ackyLog"
)

// Test deserialize all formats.
func TestMarshalling(t *testing.T) {

	jsonString := `{
	"AEAJM": {
	  "name": "Ajman",
	  "city": "Ajman",
	  "country": "United Arab Emirates",
	  "alias": [],
	  "regions": [],
	  "coordinates": [
		55.5136433,
		25.4052165
	  ],
	  "province": "Ajman",
	  "timezone": "Asia/Dubai",
	  "unlocs": [
		"AEAJM"
	  ],
	  "code": "52000"
	}
	}`

	d, err := UnMarshalWholeMap([]byte(jsonString))
	if err != nil {
		t.Errorf("Unmarshaling Failed %v", err)
	}

	for k, v := range d {
		l.INFO("PortName: %s  %+v", k, v)
	}

	//P.Unlocs = []string{"UNLOCS"}
	//P.Timezone = "TIMEZONE"
	//P.Regions = []string{"Regions"}
	//P.Province = "PROVINCE"
	//P.Name = "NAME"
	//P.Country = "COUNTRY"
	//P.Code = "CODE"
	//P.Alias = []string{"Alias"}
	//P.City = "CITY"
	//P.Coordinates = []float64{0.3434, 0.565}

}
