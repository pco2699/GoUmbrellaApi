package Server

import (
	"encoding/json"
	"log"
	"net/http"
	"umbrellaApi/WeatherApiClient"
)

type Server struct {
	ApiClient WeatherApiClient.ApiClient
}

func (s *Server) Start() {
	http.HandleFunc("/v1", s.handleApi)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func (s *Server) handleApi(w http.ResponseWriter, r *http.Request) {
	wd := s.ApiClient.Fetch()
	rainVolume := 0.0
	for _, values := range wd.List {
		rainVolume += values.Rain.Volume
	}

	isUmbrellaRequired := false
	if rainVolume > 0 {
		isUmbrellaRequired = true
	}
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"is_umbrellaRequired": isUmbrellaRequired, "rain_volume": rainVolume}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
