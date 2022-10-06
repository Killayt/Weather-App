package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Killayt/Weather-App/config"
)

type WeatherDate struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func hello(w http.ResponseWriter, r http.Request) {
	w.Write([]byte("Hello"))
}

func query(city string) (WeatherDate, error) {
	apiConfig, err := config.LoadApiConfig(".apiConfig")
	if err != nil {
		return WeatherDate{}, err
	}

	lang := "{rus}" // Choose ur language
	http.Get("https://api.openweathermap.org/data/3.0/onecall?lat={lat}&lon={lon}&lang= " + lang + "&exclude={part}&appid=" + apiConfig.ApiKey)
}

func main() {
	const port string = "9000"
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]

			date, err := query(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(date)
		})

	http.ListenAndServe(":"+port, nil)
}
