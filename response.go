package go_open_meteo

import (
	"math"
	"time"
)

type ForecastResponse struct {
	Latitude             float64             `json:"latitude"`
	Longitude            float64             `json:"longitude"`
	Elevation            float64             `json:"elevation"`
	Timezone             string              `json:"timezone"`
	GenerationTimeMs     float64             `json:"generationtime_ms"`
	UTCOffset            float64             `json:"utc_offset"`
	TimezoneAbbreviation string              `json:"timezone_abbreviation"`
	Hourly               *HourlyResponse     `json:"hourly"`
	Minutely15           *Minutely15Response `json:"minutely_15"`
	Daily                *DailyResponse      `json:"daily"`
}

type HourlyResponse struct {
	Time                     []CustomTime          `json:"time"`
	Temperature2m            []float64             `json:"temperature_2m"`
	RelativeHumidity2m       []float64             `json:"relative_humidity_2m"`
	DewPoint2m               []float64             `json:"dew_point_2m"`
	ApparentTemperature      []float64             `json:"apparent_temperature"`
	PressureMSL              []float64             `json:"pressure_msl"`
	SurfacePressure          []float64             `json:"surface_pressure"`
	CloudCover               []float64             `json:"cloud_cover"`
	CloudCoverLow            []float64             `json:"cloud_cover_low"`
	CloudCoverMid            []float64             `json:"cloud_cover_mid"`
	CloudCoverHigh           []float64             `json:"cloud_cover_high"`
	WindSpeed10m             []float64             `json:"wind_speed_10m"`
	WindSpeed80m             []float64             `json:"wind_speed_80m"`
	WindSpeed120m            []float64             `json:"wind_speed_120m"`
	WindSpeed180m            []float64             `json:"wind_speed_180m"`
	WindDirection10m         []float64             `json:"wind_direction_10m"`
	WindDirection80m         []float64             `json:"wind_direction_80m"`
	WindDirection120m        []float64             `json:"wind_direction_120m"`
	WindDirection180m        []float64             `json:"wind_direction_180m"`
	WindGusts10m             []float64             `json:"wind_gusts_10m"`
	ShortwaveRadiation       []float64             `json:"shortwave_radiation"`
	DirectRadiation          []float64             `json:"direct_radiation"`
	DirectNormalIrradiance   []float64             `json:"direct_normal_irradiance"`
	DiffuseRadiation         []float64             `json:"diffuse_radiation"`
	GlobalTiltedIrradiance   []float64             `json:"global_tilted_irradiance"`
	VapourPressureDeficit    []float64             `json:"vapour_pressure_deficit"`
	Cape                     []float64             `json:"cape"`
	Evapotranspiration       []float64             `json:"evapotranspiration"`
	Et0FaoEvapotranspiration []float64             `json:"et0_fao_evapotranspiration"`
	Precipitation            []float64             `json:"precipitation"`
	Snowfall                 []float64             `json:"snowfall"`
	PrecipitationProbability []float64             `json:"precipitation_probability"`
	Rain                     []float64             `json:"rain"`
	Showers                  []float64             `json:"showers"`
	WeatherCode              []WeatherCodeResponse `json:"weather_code"`
	SnowDepth                []float64             `json:"snow_depth"`
	FreezingLevelHeight      []float64             `json:"freezing_level_height"`
	Visibility               []float64             `json:"visibility"`
	SoilTemperature0cm       []float64             `json:"soil_temperature_0cm"`
	SoilTemperature6cm       []float64             `json:"soil_temperature_6cm"`
	SoilTemperature18cm      []float64             `json:"soil_temperature_18cm"`
	SoilTemperature54cm      []float64             `json:"soil_temperature_54cm"`
	SoilMoisture0To1cm       []float64             `json:"soil_moisture_0_to_1cm"`
	SoilMoisture1To3cm       []float64             `json:"soil_moisture_1_to_3cm"`
	SoilMoisture3To9cm       []float64             `json:"soil_moisture_3_to_9cm"`
	SoilMoisture9To27cm      []float64             `json:"soil_moisture_9_to_27cm"`
	SoilMoisture27To81cm     []float64             `json:"soil_moisture_27_to_81cm"`
	IsDay                    []int                 `json:"is_day"`
}

type Minutely15Response struct {
	Time                          []CustomTime          `json:"time"`
	Temperature2m                 []float64             `json:"temperature_2m"`
	RelativeHumidity2m            []float64             `json:"relative_humidity_2m"`
	DewPoint2m                    []float64             `json:"dew_point_2m"`
	ApparentTemperature           []float64             `json:"apparent_temperature"`
	ShortwaveRadiation            []float64             `json:"shortwave_radiation"`
	DirectRadiation               []float64             `json:"direct_radiation"`
	DirectNormalIrradiance        []float64             `json:"direct_normal_irradiance"`
	GlobalTiltedIrradiance        []float64             `json:"global_tilted_irradiance"`
	GlobalTiltedIrradianceInstant []float64             `json:"global_tilted_irradiance_instant"`
	DiffuseRadiation              []float64             `json:"diffuse_radiation"`
	SunshineDuration              []float64             `json:"sunshine_duration"`
	LightningPotential            []float64             `json:"lightning_potential"`
	Precipitation                 []float64             `json:"precipitation"`
	Snowfall                      []float64             `json:"snowfall"`
	Rain                          []float64             `json:"rain"`
	Showers                       []float64             `json:"showers"`
	SnowfallHeight                []float64             `json:"snowfall_height"`
	FreezingLevelHeight           []float64             `json:"freezing_level_height"`
	Cape                          []float64             `json:"cape"`
	WindSpeed10m                  []float64             `json:"wind_speed_10m"`
	WindSpeed80m                  []float64             `json:"wind_speed_80m"`
	WindDirection10m              []float64             `json:"wind_direction_10m"`
	WindDirection80m              []float64             `json:"wind_direction_80m"`
	WindGusts10m                  []float64             `json:"wind_gusts_10m"`
	Visibility                    []float64             `json:"visibility"`
	WeatherCode                   []WeatherCodeResponse `json:"weather_code"`
}

type DailyResponse struct {
	Time                         []CustomTime          `json:"time"`
	Temperature2mMax             []float64             `json:"temperature_2m_max"`
	Temperature2mMin             []float64             `json:"temperature_2m_min"`
	ApparentTemperatureMax       []float64             `json:"apparent_temperature_max"`
	ApparentTemperatureMin       []float64             `json:"apparent_temperature_min"`
	PrecipitationSum             []float64             `json:"precipitation_sum"`
	RainSum                      []float64             `json:"rain_sum"`
	ShowersSum                   []float64             `json:"showers_sum"`
	SnowfallSum                  []float64             `json:"snowfall_sum"`
	PrecipitationHours           []float64             `json:"precipitation_hours"`
	PrecipitationProbabilityMax  []float64             `json:"precipitation_probability_max"`
	PrecipitationProbabilityMin  []float64             `json:"precipitation_probability_min"`
	PrecipitationProbabilityMean []float64             `json:"precipitation_probability_mean"`
	WeatherCode                  []WeatherCodeResponse `json:"weather_code"`
	Sunrise                      []string              `json:"sunrise"`
	Sunset                       []string              `json:"sunset"`
	SunshineDuration             []float64             `json:"sunshine_duration"`
	DaylightDuration             []float64             `json:"daylight_duration"`
	WindSpeed10mMax              []float64             `json:"wind_speed_10m_max"`
	WindGusts10mMax              []float64             `json:"wind_gusts_10m_max"`
	WindDirection10mDominant     []float64             `json:"wind_direction_10m_dominant"`
	ShortwaveRadiationSum        []float64             `json:"shortwave_radiation_sum"`
	Et0FaoEvapotranspiration     []float64             `json:"et0_fao_evapotranspiration"`
	UvIndexMax                   []float64             `json:"uv_index_max"`
	UvIndexClearSkyMax           []float64             `json:"uv_index_clear_sky_max"`
}

func (f *ForecastResponse) FindNearestHourlyResponse(dateTime time.Time) *NearestHourlyForecast {
	hourly := f.Hourly
	closestIndex := findClosestIndex(hourly.Time, dateTime)

	return &NearestHourlyForecast{
		Time:                     hourly.Time[closestIndex],
		Temperature2m:            safeGetFloat64(hourly.Temperature2m, closestIndex),
		RelativeHumidity2m:       safeGetFloat64(hourly.RelativeHumidity2m, closestIndex),
		DewPoint2m:               safeGetFloat64(hourly.DewPoint2m, closestIndex),
		ApparentTemperature:      safeGetFloat64(hourly.ApparentTemperature, closestIndex),
		PressureMSL:              safeGetFloat64(hourly.PressureMSL, closestIndex),
		SurfacePressure:          safeGetFloat64(hourly.SurfacePressure, closestIndex),
		CloudCover:               safeGetFloat64(hourly.CloudCover, closestIndex),
		CloudCoverLow:            safeGetFloat64(hourly.CloudCoverLow, closestIndex),
		CloudCoverMid:            safeGetFloat64(hourly.CloudCoverMid, closestIndex),
		CloudCoverHigh:           safeGetFloat64(hourly.CloudCoverHigh, closestIndex),
		WindSpeed10m:             safeGetFloat64(hourly.WindSpeed10m, closestIndex),
		WindSpeed80m:             safeGetFloat64(hourly.WindSpeed80m, closestIndex),
		WindSpeed120m:            safeGetFloat64(hourly.WindSpeed120m, closestIndex),
		WindSpeed180m:            safeGetFloat64(hourly.WindSpeed180m, closestIndex),
		WindDirection10m:         safeGetFloat64(hourly.WindDirection10m, closestIndex),
		WindDirection80m:         safeGetFloat64(hourly.WindDirection80m, closestIndex),
		WindDirection120m:        safeGetFloat64(hourly.WindDirection120m, closestIndex),
		WindDirection180m:        safeGetFloat64(hourly.WindDirection180m, closestIndex),
		WindGusts10m:             safeGetFloat64(hourly.WindGusts10m, closestIndex),
		ShortwaveRadiation:       safeGetFloat64(hourly.ShortwaveRadiation, closestIndex),
		DirectRadiation:          safeGetFloat64(hourly.DirectRadiation, closestIndex),
		DirectNormalIrradiance:   safeGetFloat64(hourly.DirectNormalIrradiance, closestIndex),
		DiffuseRadiation:         safeGetFloat64(hourly.DiffuseRadiation, closestIndex),
		GlobalTiltedIrradiance:   safeGetFloat64(hourly.GlobalTiltedIrradiance, closestIndex),
		VapourPressureDeficit:    safeGetFloat64(hourly.VapourPressureDeficit, closestIndex),
		Cape:                     safeGetFloat64(hourly.Cape, closestIndex),
		Evapotranspiration:       safeGetFloat64(hourly.Evapotranspiration, closestIndex),
		Et0FaoEvapotranspiration: safeGetFloat64(hourly.Et0FaoEvapotranspiration, closestIndex),
		Precipitation:            safeGetFloat64(hourly.Precipitation, closestIndex),
		Snowfall:                 safeGetFloat64(hourly.Snowfall, closestIndex),
		PrecipitationProbability: safeGetFloat64(hourly.PrecipitationProbability, closestIndex),
		Rain:                     safeGetFloat64(hourly.Rain, closestIndex),
		Showers:                  safeGetFloat64(hourly.Showers, closestIndex),
		WeatherCode:              safeGetWeatherCode(hourly.WeatherCode, closestIndex),
		SnowDepth:                safeGetFloat64(hourly.SnowDepth, closestIndex),
		FreezingLevelHeight:      safeGetFloat64(hourly.FreezingLevelHeight, closestIndex),
		Visibility:               safeGetFloat64(hourly.Visibility, closestIndex),
		SoilTemperature0cm:       safeGetFloat64(hourly.SoilTemperature0cm, closestIndex),
		SoilTemperature6cm:       safeGetFloat64(hourly.SoilTemperature6cm, closestIndex),
		SoilTemperature18cm:      safeGetFloat64(hourly.SoilTemperature18cm, closestIndex),
		SoilTemperature54cm:      safeGetFloat64(hourly.SoilTemperature54cm, closestIndex),
		SoilMoisture0To1cm:       safeGetFloat64(hourly.SoilMoisture0To1cm, closestIndex),
		SoilMoisture1To3cm:       safeGetFloat64(hourly.SoilMoisture1To3cm, closestIndex),
		SoilMoisture3To9cm:       safeGetFloat64(hourly.SoilMoisture3To9cm, closestIndex),
		SoilMoisture9To27cm:      safeGetFloat64(hourly.SoilMoisture9To27cm, closestIndex),
		SoilMoisture27To81cm:     safeGetFloat64(hourly.SoilMoisture27To81cm, closestIndex),
		IsDay:                    safeGetInt(hourly.IsDay, closestIndex),
	}
}

func (f *ForecastResponse) FindNearestMinute15Response(dateTime time.Time) *NearestMinute15Forecast {
	minutely15 := f.Minutely15
	closestIndex := findClosestIndex(minutely15.Time, dateTime)

	return &NearestMinute15Forecast{
		Time:                          minutely15.Time[closestIndex],
		Temperature2m:                 safeGetFloat64(minutely15.Temperature2m, closestIndex),
		RelativeHumidity2m:            safeGetFloat64(minutely15.RelativeHumidity2m, closestIndex),
		DewPoint2m:                    safeGetFloat64(minutely15.DewPoint2m, closestIndex),
		ApparentTemperature:           safeGetFloat64(minutely15.ApparentTemperature, closestIndex),
		ShortwaveRadiation:            safeGetFloat64(minutely15.ShortwaveRadiation, closestIndex),
		DirectRadiation:               safeGetFloat64(minutely15.DirectRadiation, closestIndex),
		DirectNormalIrradiance:        safeGetFloat64(minutely15.DirectNormalIrradiance, closestIndex),
		GlobalTiltedIrradiance:        safeGetFloat64(minutely15.GlobalTiltedIrradiance, closestIndex),
		GlobalTiltedIrradianceInstant: safeGetFloat64(minutely15.GlobalTiltedIrradianceInstant, closestIndex),
		DiffuseRadiation:              safeGetFloat64(minutely15.DiffuseRadiation, closestIndex),
		SunshineDuration:              safeGetFloat64(minutely15.SunshineDuration, closestIndex),
		LightningPotential:            safeGetFloat64(minutely15.LightningPotential, closestIndex),
		Precipitation:                 safeGetFloat64(minutely15.Precipitation, closestIndex),
		Snowfall:                      safeGetFloat64(minutely15.Snowfall, closestIndex),
		Rain:                          safeGetFloat64(minutely15.Rain, closestIndex),
		Showers:                       safeGetFloat64(minutely15.Showers, closestIndex),
		SnowfallHeight:                safeGetFloat64(minutely15.SnowfallHeight, closestIndex),
		FreezingLevelHeight:           safeGetFloat64(minutely15.FreezingLevelHeight, closestIndex),
		Cape:                          safeGetFloat64(minutely15.Cape, closestIndex),
		WindSpeed10m:                  safeGetFloat64(minutely15.WindSpeed10m, closestIndex),
		WindSpeed80m:                  safeGetFloat64(minutely15.WindSpeed80m, closestIndex),
		WindDirection10m:              safeGetFloat64(minutely15.WindDirection10m, closestIndex),
		WindDirection80m:              safeGetFloat64(minutely15.WindDirection80m, closestIndex),
		WindGusts10m:                  safeGetFloat64(minutely15.WindGusts10m, closestIndex),
		Visibility:                    safeGetFloat64(minutely15.Visibility, closestIndex),
		WeatherCode:                   safeGetWeatherCode(minutely15.WeatherCode, closestIndex),
	}
}

func (f *ForecastResponse) FindNearestDailyResponse(dateTime time.Time) *NearestDailyForecast {
	daily := f.Daily
	closestIndex := findClosestIndex(daily.Time, dateTime)

	return &NearestDailyForecast{
		Time:                         daily.Time[closestIndex],
		Temperature2mMax:             safeGetFloat64(daily.Temperature2mMax, closestIndex),
		Temperature2mMin:             safeGetFloat64(daily.Temperature2mMin, closestIndex),
		ApparentTemperatureMax:       safeGetFloat64(daily.ApparentTemperatureMax, closestIndex),
		ApparentTemperatureMin:       safeGetFloat64(daily.ApparentTemperatureMin, closestIndex),
		PrecipitationSum:             safeGetFloat64(daily.PrecipitationSum, closestIndex),
		RainSum:                      safeGetFloat64(daily.RainSum, closestIndex),
		ShowersSum:                   safeGetFloat64(daily.ShowersSum, closestIndex),
		SnowfallSum:                  safeGetFloat64(daily.SnowfallSum, closestIndex),
		PrecipitationHours:           safeGetFloat64(daily.PrecipitationHours, closestIndex),
		PrecipitationProbabilityMax:  safeGetFloat64(daily.PrecipitationProbabilityMax, closestIndex),
		PrecipitationProbabilityMin:  safeGetFloat64(daily.PrecipitationProbabilityMin, closestIndex),
		PrecipitationProbabilityMean: safeGetFloat64(daily.PrecipitationProbabilityMean, closestIndex),
		WeatherCode:                  safeGetWeatherCode(daily.WeatherCode, closestIndex),
		Sunrise:                      safeGetString(daily.Sunrise, closestIndex),
		Sunset:                       safeGetString(daily.Sunset, closestIndex),
		SunshineDuration:             safeGetFloat64(daily.SunshineDuration, closestIndex),
		DaylightDuration:             safeGetFloat64(daily.DaylightDuration, closestIndex),
		WindSpeed10mMax:              safeGetFloat64(daily.WindSpeed10mMax, closestIndex),
		WindGusts10mMax:              safeGetFloat64(daily.WindGusts10mMax, closestIndex),
		WindDirection10mDominant:     safeGetFloat64(daily.WindDirection10mDominant, closestIndex),
		ShortwaveRadiationSum:        safeGetFloat64(daily.ShortwaveRadiationSum, closestIndex),
		Et0FaoEvapotranspiration:     safeGetFloat64(daily.Et0FaoEvapotranspiration, closestIndex),
		UvIndexMax:                   safeGetFloat64(daily.UvIndexMax, closestIndex),
		UvIndexClearSkyMax:           safeGetFloat64(daily.UvIndexClearSkyMax, closestIndex),
	}
}

func findClosestIndex(times []CustomTime, target time.Time) int {
	var closestIndex int
	minDiff := time.Duration(math.MaxInt64)

	for i, t := range times {
		diff := math.Abs(float64(target.Sub(t.Time)))
		if time.Duration(diff) < minDiff {
			minDiff = time.Duration(diff)
			closestIndex = i
		}
	}

	return closestIndex
}

const customTimeFormat = "2006-01-02T15:04"

// UnmarshalJSON method to parse the custom time format
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// Remove the surrounding quotes from the JSON string
	timeString := string(b)
	timeString = timeString[1 : len(timeString)-1]

	// Parse the time string using the custom format
	parsedTime, err := time.Parse(customTimeFormat, timeString)
	if err != nil {
		return err
	}

	ct.Time = parsedTime
	return nil
}

type CustomTime struct {
	time.Time
}

type WeatherCodeResponse int // WeatherCode type

const (
	WeatherCodeClearSky             WeatherCodeResponse = 0
	WeatherCodeMainlyClear          WeatherCodeResponse = 1
	WeatherCodePartlyCloudy         WeatherCodeResponse = 2
	WeatherCodeOvercast             WeatherCodeResponse = 3
	WeatherCodeFog                  WeatherCodeResponse = 45
	WeatherCodeDepositingRimeFog    WeatherCodeResponse = 48
	WeatherCodeLightDrizzle         WeatherCodeResponse = 51
	WeatherCodeModerateDrizzle      WeatherCodeResponse = 53
	WeatherCodeDenseDrizzle         WeatherCodeResponse = 55
	WeatherCodeLightFreezingDrizzle WeatherCodeResponse = 56
	WeatherCodeDenseFreezingDrizzle WeatherCodeResponse = 57
	WeatherCodeSlightRain           WeatherCodeResponse = 61
	WeatherCodeModerateRain         WeatherCodeResponse = 63
	WeatherCodeHeavyRain            WeatherCodeResponse = 65
	WeatherCodeLightFreezingRain    WeatherCodeResponse = 66
	WeatherCodeHeavyFreezingRain    WeatherCodeResponse = 67
	WeatherCodeSlightSnowFall       WeatherCodeResponse = 71
	WeatherCodeModerateSnowFall     WeatherCodeResponse = 73
	WeatherCodeHeavySnowFall        WeatherCodeResponse = 75
	WeatherCodeSnowGrains           WeatherCodeResponse = 77
	WeatherCodeSlightRainShowers    WeatherCodeResponse = 80
	WeatherCodeModerateRainShowers  WeatherCodeResponse = 81
	WeatherCodeViolentRainShowers   WeatherCodeResponse = 82
	WeatherCodeSlightSnowShowers    WeatherCodeResponse = 85
	WeatherCodeHeavySnowShowers     WeatherCodeResponse = 86
	WeatherCodeThunderstorm         WeatherCodeResponse = 95
	WeatherCodeSlightHailThunder    WeatherCodeResponse = 96
	WeatherCodeHeavyHailThunder     WeatherCodeResponse = 99
)
