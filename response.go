package go_open_meteo

import "time"

type Response struct {
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
