package Server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"umbrellaApi/WeatherData"
)

type testApiClient struct{}

func (t testApiClient) Fetch() WeatherData.WeatherData {
	wd := WeatherData.WeatherData{
		// テストデータを詰める
	}
	return wd
}

type ApiData struct {
	IsUmbrellaRequired bool    `json:"is_umbrella_required"`
	RainVolume         float64 `json:"rain_volume"`
}

func TestServer(t *testing.T) {
	server := &Server{
		testApiClient{},
	}
	ts := httptest.NewServer(http.HandlerFunc(server.handleApi))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("%v", err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("%v", err)
	}
	var apiData ApiData
	err = json.Unmarshal(data, &apiData)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if apiData.IsUmbrellaRequired != true {
		t.Errorf("got: %v\nwant: %v", apiData.IsUmbrellaRequired, true)
	}

	if apiData.RainVolume != 0 {
		t.Errorf("got: %v\nwant: %v", apiData.RainVolume, true)
	}

}
