package routing

import (
	"encoding/json"
	"strconv"
	"team-project/services/data"
	"testing"
)

//test if path from Lviv to Kyiv contains TrainsId
func testRouting(t *testing.T) {

	cityJSON := `{
	"start":"Lviv"",
	"end":"Kyiv"
}`
	var (
		stations   data.Stations
		pathResult []data.Routes
		sum        int
	)

	if initialised == false {
		RouteStorage = getData()
	}

	json.Unmarshal([]byte(cityJSON), &stations)

	for key, value := range RouteStorage {
		if indexStart := funk.IndexOf(value, stations.StartRoute); indexStart != -1 {
			if indexEnd := funk.IndexOf(value, stations.EndRoute); indexEnd != -1 {
				result := data.Routes{key, stations}
				pathResult = append(pathResult, result)
			}

		}
	}

	for _, v := range pathResult {
		conv, _ := strconv.Atoi(v.RouteID)
		sum += conv
	}

	if sum != 1848 {
		t.Error("Router not working")
	}

}