package main

import (
	"net/http"

	"github.com/Killayt/Weather-App/config"
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
	const port string = "9000"

	http.HandleFunc("/", Check)
	http.HandleFunc("/weather/", config.Weather)
	http.ListenAndServe(":"+port, nil)
}
