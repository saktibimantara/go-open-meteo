package go_open_meteo

type Config struct {
	APIKey        string
	BaseURL       string
	ForecastURL   string
	AQIBaseURL    string
	AirQualityURL string
}

type IConfig interface {
	GetForecastURL() string
	GetAirQualityURL() string
}

func NewConfig() *Config {
	return &Config{
		APIKey:        "",
		BaseURL:       "https://api.open-meteo.com",
		ForecastURL:   "/v1/forecast",
		AQIBaseURL:    "https://air-quality-api.open-meteo.com",
		AirQualityURL: "/v1/air-quality",
	}
}

func (c *Config) GetForecastURL() string {
	return c.BaseURL + c.ForecastURL
}

func (c *Config) GetAirQualityURL() string {
	return c.AQIBaseURL + c.AirQualityURL
}
