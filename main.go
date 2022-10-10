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
		Celsius float64 `json:"temperature"`
	} `json:"main"`
}

func check(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte("Hello!\nYes, Its works"))

	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func target(city string) (WeatherDate, error) {
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

func weather(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	date, err := target(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(date)
}

func main() {

	// GUI
	gui.MakeGUI()

	// server
	const port string = "9000"
	http.HandleFunc("/check", check)
	http.HandleFunc("/weather/", weather)
	http.ListenAndServe(":"+port, nil)

}
