package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Killayt/Weather-App/config"
	"github.com/Killayt/Weather-App/gui"
)

type WeatherDate struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

// Check connetion
func Check(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello!\nYes, it`s works!"))
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func Target(city string) (WeatherDate, error) {
	apiConfig, err := config.LoadApiConfig(".apiConfig")
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

func main() {

	// GUI

	// server
	const port string = "9000"

	go http.HandleFunc("/", Check)
	go http.HandleFunc("/weather/", Weather)
	go http.ListenAndServe(":"+port, nil)

	gui.MakeGUI()

}
