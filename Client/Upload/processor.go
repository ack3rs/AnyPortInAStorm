package Upload

import (
	"encoding/json"
	"io/ioutil"

	"github.com/acky666/AnyPortInAStorm/Client/rpc"
	pc "github.com/acky666/AnyPortInAStorm/PortController"
	l "github.com/acky666/ackyLog"
)

func processor(filename string) {

	defer l.TIMED("Processing %s", filename)()

	// TODO: This is not right, come back if you have time!
	// Load the ENTIRE file into memory
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		l.ERROR("Error reading file %v", err)
	}

	// Convert it to a Map of PortInfo
	dataMap, err := pc.UnMarshalWholeMap(dat)
	if err != nil {
		l.ERROR("Unable to UnMarshal the data: %v", err)
	}

	// Itereate the Map and Call and make the RPC
	for k, v := range dataMap {

		portData, err := json.Marshal(v)
		if err != nil {
			l.ERROR("Unable to UnMarshal the data: %v", err)
		} else {
			rpc.UpdatePort(k, string(portData))
		}
	}
}
