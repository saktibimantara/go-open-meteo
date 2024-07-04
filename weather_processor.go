package go_open_meteo

import "time"

type IWeatherData interface {
	GetForecastResponse() *ForecastResponse
}

type WeatherData struct {
	forecastResp *ForecastResponse
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (w *WeatherData) SetForecastResponse(resp *ForecastResponse) *WeatherData {
	w.forecastResp = resp

	return w
}

func (w *WeatherData) GetForecastResponse() *ForecastResponse {
	return w.forecastResp
}

type WeatherProcessor struct {
	forecastResp *ForecastResponse
}

type NearestForecast struct {
	DateTime           time.Time
	HourlyForecast     *NearestHourlyForecast
	Minutely15Forecast *NearestMinute15Forecast
	DailyForecast      *NearestDailyForecast
}

func NewWeatherProcessor(data IWeatherData) *WeatherProcessor {
	return &WeatherProcessor{
		forecastResp: data.GetForecastResponse(),
	}
}

func (w *WeatherProcessor) FindNearestForecastByTime(dateTime time.Time) (*NearestForecast, error) {

	if w.forecastResp == nil {
		return nil, ErrForecastResponseNil
	}

	nearestForecast := &NearestForecast{
		DateTime: dateTime,
	}

	if w.forecastResp.Hourly != nil {
		nearestForecast.HourlyForecast = w.forecastResp.FindNearestHourlyResponse(dateTime)
	}

	if w.forecastResp.Minutely15 != nil {
		nearestForecast.Minutely15Forecast = w.forecastResp.FindNearestMinute15Response(dateTime)
	}

	if w.forecastResp.Daily != nil {
		nearestForecast.DailyForecast = w.forecastResp.FindNearestDailyResponse(dateTime)
	}

	return nearestForecast, nil
}
