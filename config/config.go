package config

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type apiConfigDate struct {
	ApiKey string `json:"ApiKey"`
}

type WeatherDate struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func LoadApiConfig(filename string) (apiConfigDate, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return apiConfigDate{}, err
	}

	var c apiConfigDate
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigDate{}, err
	}
	return c, nil
}

func Weather(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	date, err := Target(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(date)
}

func Target(city string) (WeatherDate, error) {
	apiConfig, err := LoadApiConfig(".apiConfig")
	if err != nil {
		return WeatherDate{}, err
	}

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric" + "&appid=" + apiConfig.ApiKey)
	if err != nil {
		return WeatherDate{}, err
	}

	defer resp.Body.Close()

	var d WeatherDate
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return WeatherDate{}, err
	}

	return d, nil
}
