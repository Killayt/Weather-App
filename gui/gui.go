package gui

import (
	"encoding/json"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Killayt/Weather-App/config"
)

func MakeGUI() {
	a := app.New()
	w := a.NewWindow("Weather K")

	w.Resize(fyne.Size{Width: 500, Height: 700})

	// FIND FORM
	input := widget.NewEntry()
	input.SetPlaceHolder("City")

	content := container.NewVBox(input, widget.NewButton("Find", func() {
		log.Println("Content was", input.Text)
	}))

	w.SetContent(content)

	// MAIN FORM

	// w.SetContent(widget.NewButton("Find"))

	w.ShowAndRun()
}

func getWeatherData(city string) (config.WeatherDate, error) {
	apiConfig, err := config.LoadApiConfig(".apiConfig")
	if err != nil {
		return config.WeatherDate{}, err
	}

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric" + "&appid=" + apiConfig.ApiKey)
	if err != nil {
		return config.WeatherDate{}, err
	}

	defer resp.Body.Close()

	var d config.WeatherDate
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return config.WeatherDate{}, err
	}

	return d, nil
}

// func getWeatherData(city string) (*Target) {
// 	config.Target(city string) (WeatherDate, error) {
// 		apiConfig, err := LoadApiConfig(".apiConfig")
// 		if err != nil {
// 			return WeatherDate{}, err
// 		}

// 		resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric" + "&appid=" + apiConfig.ApiKey)
// 		if err != nil {
// 			return WeatherDate{}, err
// 		}

// 		defer resp.Body.Close()

// 		var d WeatherDate
// 		if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
// 			return WeatherDate{}, err
// 		}

// 		return d, nil
// 	}
// }
