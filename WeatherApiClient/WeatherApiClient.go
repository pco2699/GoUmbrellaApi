package WeatherApiClient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"umbrellaApi/WeatherData"
)

type WeatherApiClient struct {
	RequestUrl string
	Apikey     string
	Place      string
	Format     string
	Count      string
	Lang       string
}

func NewWeatherApiClient(Apikey string) ApiClient {
	return &WeatherApiClient{
		"https://api.openweathermap.org/data/2.5/forecast",
		Apikey,
		"Tokyo,jp",
		"json",
		"8",
		"ja",
	}
}

func (w *WeatherApiClient) Fetch() WeatherData.WeatherData {
	values := url.Values{}
	values.Add("APPID", w.Apikey)
	values.Add("q", w.Place)
	values.Add("mode", w.Format)
	values.Add("cnt", w.Count)
	values.Add("lang", w.Lang)
	res, err := http.Get(w.RequestUrl + "?" + values.Encode())
	if err != nil {
		log.Fatalf("Failed to get request url. err: %v", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var wd WeatherData.WeatherData
	err = json.Unmarshal(body, &wd)
	if err != nil {
		log.Fatalf("Failed to fetched unmarshal json. err: %v", err)
	}

	return wd
}
