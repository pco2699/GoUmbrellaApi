package WeatherData

import "encoding/json"

type WeatherData struct {
	Cod     json.Number   `json:"cod"`
	Message float64       `json:"message"`
	Count   int           `json:"count"`
	List    []WeatherList `json:"list"`
}

type WeatherList struct {
	Date    int64           `json:"dt"`
	Main    WeatherMainData `json:"main"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Rain struct {
		Volume float64 `json:"3h"`
	}
	Sys struct {
		Pod string `json:"pod"`
	} `json:"sys"`
	DtText string `json:"dt_text"`
}

type WeatherMainData struct {
	Temp        float64 `json:"temp"`
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	Pressure    float64 `json:"pressure"`
	SeaLevel    float64 `json:"sea_level"`
	GroundLevel float64 `json:"grnd_level"`
	Humidity    int     `json:"humidity"`
	TempKf      float32 `json:"temp_kf"`
}
