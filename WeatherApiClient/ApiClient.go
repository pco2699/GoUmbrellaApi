package WeatherApiClient

import "umbrellaApi/WeatherData"

type ApiClient interface {
	Fetch() WeatherData.WeatherData
}
