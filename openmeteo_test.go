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
		AddHourlyParam(Temperature2m, WindSpeed10m, Precipitation, Rain, WeatherCode).
		AddMinutely15Param(Minutely15Temperature2m, Minutely15Precipitation, Minutely15Rain, Minutely15WindDirection10m, Minutely15WeatherCode).
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
