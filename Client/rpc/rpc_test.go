package rpc

import (
	"testing"
)

func TestUpdatePort(t *testing.T) {

	test1 := `{ "PORTNAME": { "name": "Port Name", "city": "Port City", "country": "The Country where the Port is", "alias": [],"regions": [], "coordinates": [ 52.5136433, 1.4052165 ], "province": "a big area", "timezone": "GMT", "unlocs": [ "PORTNAME" ], "code": "PORTCODE" }}"`

	for i := 0; i < 10; i++ {
		UpdatePort("PORTNAME", test1)
	}

}

func TestGetPort(t *testing.T) {

	for i := 0; i < 10; i++ {
		_, err := GetPort("PORTNAME")
		if err != nil {
			t.Errorf("Error: %v", err)
		}

	}

}
