package main

import (
	"net/http"

	"github.com/Killayt/Weather-App/config"
	"github.com/Killayt/Weather-App/gui"
)

// Check connetion
func Check(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello!\nYes, it`s works!"))
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func main() {
	// server
	var Port string = "9000"

	go http.HandleFunc("/", Check)
	go http.HandleFunc("/weather/", config.Weather)
	go http.ListenAndServe(":"+Port, nil)

	// gui
	gui.MakeGUI()

}
