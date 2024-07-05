package go_open_meteo

import "time"

type IWeatherData interface {
	GetForecastResponse() *ForecastResponse
	GetAQIForecastResponse() *AQIResponse
}

type WeatherData struct {
	forecastResp    *ForecastResponse
	aqiForecastResp *AQIResponse
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (w *WeatherData) SetForecastResponse(resp *ForecastResponse) *WeatherData {
	w.forecastResp = resp

	return w
}

func (w *WeatherData) SetAQIForecastResponse(resp *AQIResponse) *WeatherData {
	w.aqiForecastResp = resp

	return w
}

func (w *WeatherData) GetForecastResponse() *ForecastResponse {
	return w.forecastResp
}

func (w *WeatherData) GetAQIForecastResponse() *AQIResponse {
	return w.aqiForecastResp
}

type WeatherProcessor struct {
	forecastResp *ForecastResponse
	aqiResp      *AQIResponse
}

type NearestForecast struct {
	DateTime           time.Time
	HourlyForecast     *NearestHourlyForecast
	Minutely15Forecast *NearestMinute15Forecast
	DailyForecast      *NearestDailyForecast
	aqiHourlyForecast  *NearestAQIHourlyForecast
}

func NewWeatherProcessor(data IWeatherData) *WeatherProcessor {
	return &WeatherProcessor{
		forecastResp: data.GetForecastResponse(),
		aqiResp:      data.GetAQIForecastResponse(),
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

func (w *WeatherProcessor) FindNearestAQIForecastByTime(dateTime time.Time) (*NearestForecast, error) {
	if w.aqiResp == nil {
		return nil, ErrForecastResponseNil
	}

	nearestForecast := &NearestForecast{
		DateTime: dateTime,
	}

	if w.aqiResp.Hourly != nil {
		nearestForecast.aqiHourlyForecast = w.aqiResp.FindNearestAQIHourlyResponse(dateTime)
	}

	return nearestForecast, nil
}
