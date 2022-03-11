package backend

import (
	"fmt"
	"strings"
)

func FormatAsPlainText(info WeatherInfo) (text string) {
	var body []string
	body = append(body, "🌦️ Weather Now\n\n")

	body = append(body, fmt.Sprintf("🌍 Location: %s\n", info.location()))
	// if len(info.Kind) > 0 {
	// 	body = append(body, fmt.Sprintf("Weather description: %s\n", info.Kind[0].Description))
	// }
	body = append(body, fmt.Sprintf("🌡️ Temperature: %.2f °C\n", info.Main.Temperature))
	body = append(body, fmt.Sprintf("🌡️ Temperature feels like: %.2f °C\n", info.Main.TemperatureFeelsLike))
	body = append(body, fmt.Sprintf("⬇️ Temperature min: %.2f °C\n", info.Main.TemperatureMin))
	body = append(body, fmt.Sprintf("⬆️ Temperature max: %.2f °C\n", info.Main.TemperatureMax))
	// body = append(body, fmt.Sprintf("Wind direction: %d°\n", info.Wind.Direction))
	body = append(body, fmt.Sprintf("🌬️ Wind speed: %.2f km/h\n", convertToKmPerH(info.Wind.Speed)))
	// body = append(body, fmt.Sprintf(
	// 	"Wind gust speed: %.2f km/h\n",
	// 	convertToKmPerH(info.Wind.GustSpeed),
	// ))
	// body = append(body, fmt.Sprintf("Visibility: %d m\n", info.Visibility))
	body = append(body, fmt.Sprintf("🌧️ Precipitation volume: %.2f mm/h\n", info.precipitationVolume()))
	// body = append(body, fmt.Sprintf("Pressure: %d hPa\n", info.Main.Pressure))
	body = append(body, fmt.Sprintf("💧 Humidity: %d %%\n", info.Main.Humidity))

	return strings.Join(body, "")
}

func convertToKmPerH(value float32) float32 {
	return value * 3.6
}
