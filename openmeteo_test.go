package go_open_meteo

import (
	"testing"
	"time"
)

func TestForecast(t *testing.T) {

	cfg := NewConfig()

	client := New(cfg)

	params, err := NewForecastParamsBuilder().
		SetLatitude(-8.68163896537287).
		SetLongitude(115.19724863873421).
		SetForecastDays(16).
		AddHourlyParam(
			Temperature2m,
			WindSpeed10m,
			Precipitation,
			Rain,
			WeatherCode,
			WindSpeed10m,
			WindDirection10m,
			WindGusts10m,
		).
		AddMinutely15Param(
			Minutely15Temperature2m,
			Minutely15Precipitation,
			Minutely15Rain,
			Minutely15WeatherCode,
			Minutely15RelativeHumidity2m,
			Minutely15WindDirection10m,
			Minutely15WindSpeed10m,
			Minutely15WindGusts10m,
			Minutely15ApparentTemperature,
		).
		AddDailyParam(
			DailyTemperature2mMax,
			DailyTemperature2mMin,
			DailyWeatherCode,
		).
		Build()

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	resp, err := client.Forecast(params)

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	wd := NewWeatherData().SetForecastResponse(resp)

	wp := NewWeatherProcessor(wd)

	nf, err := wp.FindNearestForecastByTime(time.Now())

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	t.Log(nf)

}
