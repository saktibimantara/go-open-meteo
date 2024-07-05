package go_open_meteo

import "time"

type NearestHourlyForecast struct {
	Time                     CustomTime           `json:"time"`
	Temperature2m            *float64             `json:"temperature_2m"`
	RelativeHumidity2m       *float64             `json:"relative_humidity_2m"`
	DewPoint2m               *float64             `json:"dew_point_2m"`
	ApparentTemperature      *float64             `json:"apparent_temperature"`
	PressureMSL              *float64             `json:"pressure_msl"`
	SurfacePressure          *float64             `json:"surface_pressure"`
	CloudCover               *float64             `json:"cloud_cover"`
	CloudCoverLow            *float64             `json:"cloud_cover_low"`
	CloudCoverMid            *float64             `json:"cloud_cover_mid"`
	CloudCoverHigh           *float64             `json:"cloud_cover_high"`
	WindSpeed80m             *float64             `json:"wind_speed_80m"`
	WindSpeed120m            *float64             `json:"wind_speed_120m"`
	WindSpeed10m             *float64             `json:"wind_speed_10m"`
	WindSpeed180m            *float64             `json:"wind_speed_180m"`
	WindDirection10m         *float64             `json:"wind_direction_10m"`
	WindDirection80m         *float64             `json:"wind_direction_80m"`
	WindDirection120m        *float64             `json:"wind_direction_120m"`
	WindDirection180m        *float64             `json:"wind_direction_180m"`
	WindGusts10m             *float64             `json:"wind_gusts_10m"`
	ShortwaveRadiation       *float64             `json:"shortwave_radiation"`
	DirectRadiation          *float64             `json:"direct_radiation"`
	DirectNormalIrradiance   *float64             `json:"direct_normal_irradiance"`
	DiffuseRadiation         *float64             `json:"diffuse_radiation"`
	GlobalTiltedIrradiance   *float64             `json:"global_tilted_irradiance"`
	VapourPressureDeficit    *float64             `json:"vapour_pressure_deficit"`
	Cape                     *float64             `json:"cape"`
	Evapotranspiration       *float64             `json:"evapotranspiration"`
	Et0FaoEvapotranspiration *float64             `json:"et0_fao_evapotranspiration"`
	Precipitation            *float64             `json:"precipitation"`
	Snowfall                 *float64             `json:"snowfall"`
	PrecipitationProbability *float64             `json:"precipitation_probability"`
	Rain                     *float64             `json:"rain"`
	Showers                  *float64             `json:"showers"`
	WeatherCode              *WeatherCodeResponse `json:"weather_code"`
	SnowDepth                *float64             `json:"snow_depth"`
	FreezingLevelHeight      *float64             `json:"freezing_level_height"`
	Visibility               *float64             `json:"visibility"`
	SoilTemperature0cm       *float64             `json:"soil_temperature_0cm"`
	SoilTemperature6cm       *float64             `json:"soil_temperature_6cm"`
	SoilTemperature18cm      *float64             `json:"soil_temperature_18cm"`
	SoilTemperature54cm      *float64             `json:"soil_temperature_54cm"`
	SoilMoisture0To1cm       *float64             `json:"soil_moisture_0_to_1cm"`
	SoilMoisture1To3cm       *float64             `json:"soil_moisture_1_to_3cm"`
	SoilMoisture3To9cm       *float64             `json:"soil_moisture_3_to_9cm"`
	SoilMoisture9To27cm      *float64             `json:"soil_moisture_9_to_27cm"`
	SoilMoisture27To81cm     *float64             `json:"soil_moisture_27_to_81cm"`
	IsDay                    *int                 `json:"is_day"`
}

type NearestAQIHourlyForecast struct {
	Time                       *time.Time `json:"timestamp"`
	PM10                       *float64   `json:"pm10"`                          // μg/m³
	PM2_5                      *float64   `json:"pm2_5"`                         // μg/m³
	CarbonMonoxide             *float64   `json:"carbon_monoxide"`               // μg/m³
	NitrogenDioxide            *float64   `json:"nitrogen_dioxide"`              // μg/m³
	SulphurDioxide             *float64   `json:"sulphur_dioxide"`               // μg/m³
	Ozone                      *float64   `json:"ozone"`                         // μg/m³
	Ammonia                    *float64   `json:"ammonia"`                       // μg/m³
	AerosolOpticalDepth        *float64   `json:"aerosol_optical_depth"`         // Dimensionless
	Dust                       *float64   `json:"dust"`                          // μg/m³
	UVIndex                    *float64   `json:"uv_index"`                      // Index
	UVIndexClearSky            *float64   `json:"uv_index_clear_sky"`            // Index
	AlderPollen                *float64   `json:"alder_pollen"`                  // Grains/m³
	BirchPollen                *float64   `json:"birch_pollen"`                  // Grains/m³
	GrassPollen                *float64   `json:"grass_pollen"`                  // Grains/m³
	MugwortPollen              *float64   `json:"mugwort_pollen"`                // Grains/m³
	OlivePollen                *float64   `json:"olive_pollen"`                  // Grains/m³
	RagweedPollen              *float64   `json:"ragweed_pollen"`                // Grains/m³
	EuropeanAQI                *float64   `json:"european_aqi"`                  // European AQI
	EuropeanAQIPM25            *float64   `json:"european_aqi_pm2_5"`            // European AQI
	EuropeanAQIPM10            *float64   `json:"european_aqi_pm10"`             // European AQI
	EuropeanAQINitrogenDioxide *float64   `json:"european_aqi_nitrogen_dioxide"` // European AQI
	EuropeanAQIOzone           *float64   `json:"european_aqi_ozone"`            // European AQI
	EuropeanAQISulphurDioxide  *float64   `json:"european_aqi_sulphur_dioxide"`  // European AQI
	USAQI                      *float64   `json:"us_aqi"`                        // U.S. AQI
	USAQIPM25                  *float64   `json:"us_aqi_pm2_5"`                  // U.S. AQI
	USAQIPM10                  *float64   `json:"us_aqi_pm10"`                   // U.S. AQI
	USAQINitrogenDioxide       *float64   `json:"us_aqi_nitrogen_dioxide"`       // U.S. AQI
	USAQIOzone                 *float64   `json:"us_aqi_ozone"`                  // U.S. AQI
	USAQISulphurDioxide        *float64   `json:"us_aqi_sulphur_dioxide"`        // U.S. AQI
	USAQICarbonMonoxide        *float64   `json:"us_aqi_carbon_monoxide"`        // U.S. AQI
}

type NearestDailyForecast struct {
	Time                         CustomDate           `json:"time"`
	Temperature2mMax             *float64             `json:"temperature_2m_max,omitempty"`
	Temperature2mMin             *float64             `json:"temperature_2m_min,omitempty"`
	ApparentTemperatureMax       *float64             `json:"apparent_temperature_max,omitempty"`
	ApparentTemperatureMin       *float64             `json:"apparent_temperature_min,omitempty"`
	PrecipitationSum             *float64             `json:"precipitation_sum,omitempty"`
	RainSum                      *float64             `json:"rain_sum,omitempty"`
	ShowersSum                   *float64             `json:"showers_sum,omitempty"`
	SnowfallSum                  *float64             `json:"snowfall_sum,omitempty"`
	PrecipitationHours           *float64             `json:"precipitation_hours,omitempty"`
	PrecipitationProbabilityMax  *float64             `json:"precipitation_probability_max,omitempty"`
	PrecipitationProbabilityMin  *float64             `json:"precipitation_probability_min,omitempty"`
	PrecipitationProbabilityMean *float64             `json:"precipitation_probability_mean,omitempty"`
	WeatherCode                  *WeatherCodeResponse `json:"weather_code,omitempty"`
	Sunrise                      *string              `json:"sunrise,omitempty"`
	Sunset                       *string              `json:"sunset,omitempty"`
	SunshineDuration             *float64             `json:"sunshine_duration,omitempty"`
	DaylightDuration             *float64             `json:"daylight_duration,omitempty"`
	WindSpeed10mMax              *float64             `json:"wind_speed_10m_max,omitempty"`
	WindGusts10mMax              *float64             `json:"wind_gusts_10m_max,omitempty"`
	WindDirection10mDominant     *float64             `json:"wind_direction_10m_dominant,omitempty"`
	ShortwaveRadiationSum        *float64             `json:"shortwave_radiation_sum,omitempty"`
	Et0FaoEvapotranspiration     *float64             `json:"et0_fao_evapotranspiration,omitempty"`
	UvIndexMax                   *float64             `json:"uv_index_max,omitempty"`
	UvIndexClearSkyMax           *float64             `json:"uv_index_clear_sky_max,omitempty"`
}

type NearestMinute15Forecast struct {
	Time                          CustomTime           `json:"time"`
	Temperature2m                 *float64             `json:"temperature_2m,omitempty"`
	RelativeHumidity2m            *float64             `json:"relative_humidity_2m,omitempty"`
	DewPoint2m                    *float64             `json:"dew_point_2m,omitempty"`
	ApparentTemperature           *float64             `json:"apparent_temperature,omitempty"`
	ShortwaveRadiation            *float64             `json:"shortwave_radiation,omitempty"`
	DirectRadiation               *float64             `json:"direct_radiation,omitempty"`
	DirectNormalIrradiance        *float64             `json:"direct_normal_irradiance,omitempty"`
	GlobalTiltedIrradiance        *float64             `json:"global_tilted_irradiance,omitempty"`
	GlobalTiltedIrradianceInstant *float64             `json:"global_tilted_irradiance_instant,omitempty"`
	DiffuseRadiation              *float64             `json:"diffuse_radiation,omitempty"`
	SunshineDuration              *float64             `json:"sunshine_duration,omitempty"`
	LightningPotential            *float64             `json:"lightning_potential,omitempty"`
	Precipitation                 *float64             `json:"precipitation,omitempty"`
	Snowfall                      *float64             `json:"snowfall,omitempty"`
	Rain                          *float64             `json:"rain,omitempty"`
	Showers                       *float64             `json:"showers,omitempty"`
	SnowfallHeight                *float64             `json:"snowfall_height,omitempty"`
	FreezingLevelHeight           *float64             `json:"freezing_level_height,omitempty"`
	Cape                          *float64             `json:"cape,omitempty"`
	WindSpeed10m                  *float64             `json:"wind_speed_10m,omitempty"`
	WindSpeed80m                  *float64             `json:"wind_speed_80m,omitempty"`
	WindDirection10m              *float64             `json:"wind_direction_10m,omitempty"`
	WindDirection80m              *float64             `json:"wind_direction_80m,omitempty"`
	WindGusts10m                  *float64             `json:"wind_gusts_10m,omitempty"`
	Visibility                    *float64             `json:"visibility,omitempty"`
	WeatherCode                   *WeatherCodeResponse `json:"weather_code,omitempty"`
}
