package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

var weatherEmojis = map[string]string{
	"Sunny":          "☀️",
	"Clear":          "🌙",
	"Partly cloudly": "⛅",
	"Cloudy":         "☁️",
	"Overcast":       "☁️",
	"Mist":           "🌫️",
	"Patchy rain":    "🌦️",
	"Light rain":     "🌧️",
	"Moderate rain":  "🌧️",
	"Heavy rain":     "🌧️",
	"Thunderstorm":   "⛈️",
	"Snow":           "❄️",
	"Light snow":     "❄️",
	"Moderate snow":  "❄️",
	"Heavy snow":     "❄️",
	"Fog":            "🌫️",
}

func getEmoji(condition string) string {
	if emoji, ok := weatherEmojis[condition]; ok {
		return emoji
	}
	return "🌍"
}

func main() {
	godotenv.Load()
	API_KEY := os.Getenv("API_KEY")
	q := "Nairobi"
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	resp, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=" + API_KEY + "&q=" + q + "&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Weather Api not available")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	// fmt.Println(weather)
	location, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour
	fmt.Printf("%s, %s: %.0fC, %s, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
		getEmoji(current.Condition.Text),
	)
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf("%s - %.0fC, %.0f%%, %s, %s\n",
			date.Format("15:08"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
			getEmoji(hour.Condition.Text),
		)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}
