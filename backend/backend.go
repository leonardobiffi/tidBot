package backend

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetWeatherInfo() string {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("OpenWeather API Key is missing; " +
			"please, specify the OPENWEATHER_API_KEY environment variable")
	}

	city := os.Getenv("OPENWEATHER_CITY")
	if city == "" {
		log.Fatal("OpenWeather City is missing; " +
			"please, specify the OPENWEATHER_CITY environment variable")
	}

	info := WeatherInfo{}
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s&units=metric",
		apiKey,
		city,
	)

	if err := LoadJSONData(&http.Client{}, url, &info); err != nil {
		log.Fatalf("unable to load the weather data: %s", err)
	}

	return FormatAsPlainText(info)
}
