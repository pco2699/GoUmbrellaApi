package WeatherApiClient

import (
	"testing"
)

func TestWeatherApiClient_Fetch(t *testing.T) {
	wac := NewWeatherApiClient("11eccaf98a9c2e4c53520ae92d539b4f")
	wd := wac.Fetch()

	if wd.Cod.String() != "200" {
		t.Errorf("got: %v\nwant: %v", wd.Cod, "200")
	}
	for _, values := range wd.List {
		if &values.Rain.Volume == nil {
			t.Errorf("got: %v\nwant: %v", values.Rain.Volume, "Number")
		}
	}
}
