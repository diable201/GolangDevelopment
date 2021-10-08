package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	owm "github.com/briandowns/openweathermap"
)

type Weather struct {
	Temperature float64 `json:"temperature"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	Clouds      int     `json:"clouds"`
}

func GetWeather(location string) (result string) {
	os.Setenv("OWM_API_KEY", "YOUR-KEY")
	apiKey := os.Getenv("OWM_API_KEY")
	w, err := owm.NewCurrent("C", "en", apiKey)

	if err != nil {
		log.Print("getWeather() error:\n", err)
		return "error"
	}

	w.CurrentByName(location)
	result += fmt.Sprintf("Weather for 🏡 %s (%s)\n", w.Name, w.Sys.Country)
	result += fmt.Sprintf("Temperature 🌡: %.1f°C\n", w.Main.Temp)
	result += fmt.Sprintf("Humidity 💧: %d%%\n", w.Main.Humidity)
	for _, item := range w.Weather {
		result += fmt.Sprintf("%s: %s\n", item.Main, item.Description)
	}
	result += fmt.Sprintf("Wind speed 💨: %.1fm/s\n", w.Wind.Speed)
	result += fmt.Sprintf("Clouds ☁️: %d%%\n", w.Clouds.All)

	return result
}

func SerializeWeather(location string) string {
	os.Setenv("OWM_API_KEY", "YOUR-KEY")
	apiKey := os.Getenv("OWM_API_KEY")
	w, err := owm.NewCurrent("C", "en", apiKey)

	if err != nil {
		log.Print("getWeather() => newCurrent() error:\n", err)
		return "error"
	}

	w.CurrentByName(location)
	res, err := json.Marshal(Weather{Temperature: w.Main.Temp, Humidity: w.Main.Humidity,
		WindSpeed: w.Wind.Speed, Clouds: w.Clouds.All})
	if err != nil {
		panic(err)
	}
	return string(res)
}

func GetTime(location string) (result string) {
	loc, _ := time.LoadLocation(location)
	time_ := time.Now().In(loc)
	return time_.String()
}
