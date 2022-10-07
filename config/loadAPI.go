package cfg

import (
	"encoding/json"
	"os"
)

type apiConfigDate struct {
	ApiKey string `json:"ApiKey"`
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
