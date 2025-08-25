# 🌤 Weather CLI (Go)

A simple command-line tool written in **Go** that fetches the weather forecast using the [WeatherAPI](https://www.weatherapi.com/).  
It shows the **current weather** and an **hourly forecast** with a highlight (in red) when there is a high chance of rain.  

---

## 🚀 Features
- Get **current weather conditions** (temperature, location, condition).  
- **Hourly forecast** for the day.  
- Rain probability **highlighted in red** when above 40%.  
- Accepts a **city name** as a command-line argument. Defaults to `Nairobi`.  
- Uses `.env` file for secure API key storage.  

---

## 🛠️ Installation

### Prerequisites
- Go 1.22+  
- A [WeatherAPI](https://www.weatherapi.com/) key  

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/weather-cli.git
   cd weather-cli
2. Install dependencies
    go mod tidy
3. Create a .env file and add your API key:
    API_KEY=your_api_key_here
4. Run the program:
    go run main.go Nairobi
Or:without an argument (defaults to Nairobi):
go run main.go
 
