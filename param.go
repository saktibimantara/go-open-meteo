package go_open_meteo

import "fmt"

type Params struct {
	Latitude  float64
	Longitude float64
	Elevation float64
	Hourly    []string
}

type TemperatureUnit string

const (
	Celsius    TemperatureUnit = "celsius"
	Fahrenheit TemperatureUnit = "fahrenheit"
)

type IForecastParams interface {
	GetParams() string
}

type ForecastParams struct {
	latitude        float64
	longitude       float64
	forecastDays    int
	pastDays        int
	temperatureUnit TemperatureUnit
	hourlyParam     []HourlyParam
	minutely15Param []Minutely15Param
	dailyParam      []DailyParam
}

type AQIParams struct {
	latitude     float64
	longitude    float64
	forecastDays int
	pastDays     int
	hourlyParam  []AQIParam
}

func (f ForecastParams) GetParams() string {
	param := ""

	param += fmt.Sprintf("latitude=%f&longitude=%f", f.latitude, f.longitude)

	for i, p := range f.hourlyParam {
		if i == 0 {
			param += "&hourly=" + string(p)
			continue
		}

		param += "," + string(p)
	}

	for i, p := range f.minutely15Param {
		if i == 0 {
			param += "&minutely_15=" + string(p)
			continue
		}

		param += "," + string(p)
	}

	for i, p := range f.dailyParam {
		if i == 0 {
			param += "&daily=" + string(p)
			continue
		}

		param += "," + string(p)

	}

	return param
}

type ForecastParamsBuilder struct {
	forecastParams ForecastParams
}

func NewForecastParamsBuilder() *ForecastParamsBuilder {
	return &ForecastParamsBuilder{
		forecastParams: ForecastParams{},
	}
}

func (fpb *ForecastParamsBuilder) AddHourlyParam(params ...HourlyParam) *ForecastParamsBuilder {
	for _, param := range params {
		fpb.forecastParams.hourlyParam = append(fpb.forecastParams.hourlyParam, param)
	}
	return fpb
}

func (fpb *ForecastParamsBuilder) AddMinutely15Param(params ...Minutely15Param) *ForecastParamsBuilder {
	for _, param := range params {
		fpb.forecastParams.minutely15Param = append(fpb.forecastParams.minutely15Param, param)

	}
	return fpb
}

func (fpb *ForecastParamsBuilder) AddDailyParam(params ...DailyParam) *ForecastParamsBuilder {
	for _, param := range params {
		fpb.forecastParams.dailyParam = append(fpb.forecastParams.dailyParam, param)
	}
	return fpb
}

func (fpb *ForecastParamsBuilder) SetLatitude(latitude float64) *ForecastParamsBuilder {
	fpb.forecastParams.latitude = latitude
	return fpb
}

func (fpb *ForecastParamsBuilder) SetLongitude(longitude float64) *ForecastParamsBuilder {
	fpb.forecastParams.longitude = longitude
	return fpb
}

func (fpb *ForecastParamsBuilder) SetForecastDays(forecastDays int) *ForecastParamsBuilder {
	fpb.forecastParams.forecastDays = forecastDays
	return fpb
}

func (fpb *ForecastParamsBuilder) SetPastDays(pastDays int) *ForecastParamsBuilder {
	fpb.forecastParams.pastDays = pastDays
	return fpb
}

func (fpb *ForecastParamsBuilder) SetTemperatureUnit(temperatureUnit TemperatureUnit) *ForecastParamsBuilder {
	fpb.forecastParams.temperatureUnit = temperatureUnit
	return fpb
}

func (fpb *ForecastParamsBuilder) Build() (ForecastParams, error) {
	if fpb.forecastParams.latitude == 0 {
		return ForecastParams{}, fmt.Errorf("latitude is required")
	}

	if fpb.forecastParams.longitude == 0 {
		return ForecastParams{}, fmt.Errorf("longitude is required")
	}

	if fpb.forecastParams.forecastDays > 16 {
		return ForecastParams{}, fmt.Errorf("forecast days must be less than 16")
	}

	if fpb.forecastParams.pastDays > 92 {
		return ForecastParams{}, fmt.Errorf("past days must be less than 92")
	}

	return fpb.forecastParams, nil
}

type AQIParamsBuilder struct {
	aqiParams AQIParams
}

func NewAQIParamsBuilder() *AQIParamsBuilder {
	return &AQIParamsBuilder{
		aqiParams: AQIParams{},
	}
}

func (apb *AQIParamsBuilder) AddHourlyParam(params ...AQIParam) *AQIParamsBuilder {
	for _, param := range params {
		apb.aqiParams.hourlyParam = append(apb.aqiParams.hourlyParam, param)
	}
	return apb
}

func (apb *AQIParamsBuilder) SetLatitude(latitude float64) *AQIParamsBuilder {
	apb.aqiParams.latitude = latitude
	return apb
}

func (apb *AQIParamsBuilder) SetLongitude(longitude float64) *AQIParamsBuilder {
	apb.aqiParams.longitude = longitude
	return apb
}

func (apb *AQIParamsBuilder) SetForecastDays(forecastDays int) *AQIParamsBuilder {
	apb.aqiParams.forecastDays = forecastDays
	return apb
}

func (apb *AQIParamsBuilder) SetPastDays(pastDays int) *AQIParamsBuilder {
	apb.aqiParams.pastDays = pastDays
	return apb
}

func (apb *AQIParamsBuilder) Build() (AQIParams, error) {
	if apb.aqiParams.latitude == 0 {
		return AQIParams{}, fmt.Errorf("latitude is required")
	}

	if apb.aqiParams.longitude == 0 {
		return AQIParams{}, fmt.Errorf("longitude is required")
	}

	if apb.aqiParams.forecastDays > 7 {
		return AQIParams{}, fmt.Errorf("forecast days must be less than 16")
	}

	if apb.aqiParams.pastDays > 92 {
		return AQIParams{}, fmt.Errorf("past days must be less than 92")
	}

	return apb.aqiParams, nil

}

func (ap AQIParams) GetParams() string {
	param := ""

	param += fmt.Sprintf("latitude=%f&longitude=%f", ap.latitude, ap.longitude)

	for i, p := range ap.hourlyParam {
		if i == 0 {
			param += "&hourly=" + string(p)
			continue
		}

		param += "," + string(p)
	}

	if ap.forecastDays > 0 {
		param += fmt.Sprintf("&forecast_days=%d", ap.forecastDays)
	}

	if ap.pastDays > 0 {
		param += fmt.Sprintf("&past_days=%d", ap.pastDays)
	}

	return param
}

type HourlyParam string

const (
	Temperature2m            HourlyParam = "temperature_2m"
	RelativeHumidity2m       HourlyParam = "relative_humidity_2m"
	DewPoint2m               HourlyParam = "dew_point_2m"
	ApparentTemperature      HourlyParam = "apparent_temperature"
	PressureMSL              HourlyParam = "pressure_msl"
	SurfacePressure          HourlyParam = "surface_pressure"
	CloudCover               HourlyParam = "cloud_cover"
	CloudCoverLow            HourlyParam = "cloud_cover_low"
	CloudCoverMid            HourlyParam = "cloud_cover_mid"
	CloudCoverHigh           HourlyParam = "cloud_cover_high"
	WindSpeed10m             HourlyParam = "wind_speed_10m"
	WindSpeed80m             HourlyParam = "wind_speed_80m"
	WindSpeed120m            HourlyParam = "wind_speed_120m"
	WindSpeed180m            HourlyParam = "wind_speed_180m"
	WindDirection10m         HourlyParam = "wind_direction_10m"
	WindDirection80m         HourlyParam = "wind_direction_80m"
	WindDirection120m        HourlyParam = "wind_direction_120m"
	WindDirection180m        HourlyParam = "wind_direction_180m"
	WindGusts10m             HourlyParam = "wind_gusts_10m"
	ShortwaveRadiation       HourlyParam = "shortwave_radiation"
	DirectRadiation          HourlyParam = "direct_radiation"
	DirectNormalIrradiance   HourlyParam = "direct_normal_irradiance"
	DiffuseRadiation         HourlyParam = "diffuse_radiation"
	GlobalTiltedIrradiance   HourlyParam = "global_tilted_irradiance"
	VapourPressureDeficit    HourlyParam = "vapour_pressure_deficit"
	Cape                     HourlyParam = "cape"
	Evapotranspiration       HourlyParam = "evapotranspiration"
	Et0FaoEvapotranspiration HourlyParam = "et0_fao_evapotranspiration"
	Precipitation            HourlyParam = "precipitation"
	Snowfall                 HourlyParam = "snowfall"
	PrecipitationProbability HourlyParam = "precipitation_probability"
	Rain                     HourlyParam = "rain"
	Showers                  HourlyParam = "showers"
	WeatherCode              HourlyParam = "weather_code"
	SnowDepth                HourlyParam = "snow_depth"
	FreezingLevelHeight      HourlyParam = "freezing_level_height"
	Visibility               HourlyParam = "visibility"
	SoilTemperature0cm       HourlyParam = "soil_temperature_0cm"
	SoilTemperature6cm       HourlyParam = "soil_temperature_6cm"
	SoilTemperature18cm      HourlyParam = "soil_temperature_18cm"
	SoilTemperature54cm      HourlyParam = "soil_temperature_54cm"
	SoilMoisture0To1cm       HourlyParam = "soil_moisture_0_to_1cm"
	SoilMoisture1To3cm       HourlyParam = "soil_moisture_1_to_3cm"
	SoilMoisture3To9cm       HourlyParam = "soil_moisture_3_to_9cm"
	SoilMoisture9To27cm      HourlyParam = "soil_moisture_9_to_27cm"
	SoilMoisture27To81cm     HourlyParam = "soil_moisture_27_to_81cm"
	IsDay                    HourlyParam = "is_day"
)

type Minutely15Param string

const (
	Minutely15Temperature2m                 Minutely15Param = "temperature_2m"
	Minutely15RelativeHumidity2m            Minutely15Param = "relative_humidity_2m"
	Minutely15DewPoint2m                    Minutely15Param = "dew_point_2m"
	Minutely15ApparentTemperature           Minutely15Param = "apparent_temperature"
	Minutely15ShortwaveRadiation            Minutely15Param = "shortwave_radiation"
	Minutely15DirectRadiation               Minutely15Param = "direct_radiation"
	Minutely15DirectNormalIrradiance        Minutely15Param = "direct_normal_irradiance"
	Minutely15GlobalTiltedIrradiance        Minutely15Param = "global_tilted_irradiance"
	Minutely15GlobalTiltedIrradianceInstant Minutely15Param = "global_tilted_irradiance_instant"
	Minutely15DiffuseRadiation              Minutely15Param = "diffuse_radiation"
	Minutely15SunshineDuration              Minutely15Param = "sunshine_duration"
	Minutely15LightningPotential            Minutely15Param = "lightning_potential"
	Minutely15Precipitation                 Minutely15Param = "precipitation"
	Minutely15Snowfall                      Minutely15Param = "snowfall"
	Minutely15Rain                          Minutely15Param = "rain"
	Minutely15Showers                       Minutely15Param = "showers"
	Minutely15SnowfallHeight                Minutely15Param = "snowfall_height"
	Minutely15FreezingLevelHeight           Minutely15Param = "freezing_level_height"
	Minutely15Cape                          Minutely15Param = "cape"
	Minutely15WindSpeed10m                  Minutely15Param = "wind_speed_10m"
	Minutely15WindSpeed80m                  Minutely15Param = "wind_speed_80m"
	Minutely15WindDirection10m              Minutely15Param = "wind_direction_10m"
	Minutely15WindDirection80m              Minutely15Param = "wind_direction_80m"
	Minutely15WindGusts10m                  Minutely15Param = "wind_gusts_10m"
	Minutely15Visibility                    Minutely15Param = "visibility"
	Minutely15WeatherCode                   Minutely15Param = "weather_code"
)

type DailyParam string

const (
	DailyTemperature2mMax             DailyParam = "temperature_2m_max"
	DailyTemperature2mMin             DailyParam = "temperature_2m_min"
	DailyApparentTemperatureMax       DailyParam = "apparent_temperature_max"
	DailyApparentTemperatureMin       DailyParam = "apparent_temperature_min"
	DailyPrecipitationSum             DailyParam = "precipitation_sum"
	DailyRainSum                      DailyParam = "rain_sum"
	DailyShowersSum                   DailyParam = "showers_sum"
	DailySnowfallSum                  DailyParam = "snowfall_sum"
	DailyPrecipitationHours           DailyParam = "precipitation_hours"
	DailyPrecipitationProbabilityMax  DailyParam = "precipitation_probability_max"
	DailyPrecipitationProbabilityMin  DailyParam = "precipitation_probability_min"
	DailyPrecipitationProbabilityMean DailyParam = "precipitation_probability_mean"
	DailyWeatherCode                  DailyParam = "weather_code"
	DailySunrise                      DailyParam = "sunrise"
	DailySunset                       DailyParam = "sunset"
	DailySunshineDuration             DailyParam = "sunshine_duration"
	DailyDaylightDuration             DailyParam = "daylight_duration"
	DailyWindSpeed10mMax              DailyParam = "wind_speed_10m_max"
	DailyWindGusts10mMax              DailyParam = "wind_gusts_10m_max"
	DailyWindDirection10mDominant     DailyParam = "wind_direction_10m_dominant"
	DailyShortwaveRadiationSum        DailyParam = "shortwave_radiation_sum"
	DailyEt0FaoEvapotranspiration     DailyParam = "et0_fao_evapotranspiration"
	DailyUvIndexMax                   DailyParam = "uv_index_max"
	DailyUvIndexClearSkyMax           DailyParam = "uv_index_clear_sky_max"
)

type AQIParam string

const (
	PM10                       AQIParam = "pm10"
	PM2_5                      AQIParam = "pm2_5"
	CarbonMonoxide             AQIParam = "carbon_monoxide"
	NitrogenDioxide            AQIParam = "nitrogen_dioxide"
	SulphurDioxide             AQIParam = "sulphur_dioxide"
	Ozone                      AQIParam = "ozone"
	Ammonia                    AQIParam = "ammonia"
	AerosolOpticalDepth        AQIParam = "aerosol_optical_depth"
	Dust                       AQIParam = "dust"
	UVIndex                    AQIParam = "uv_index"
	UVIndexClearSky            AQIParam = "uv_index_clear_sky"
	AlderPollen                AQIParam = "alder_pollen"
	BirchPollen                AQIParam = "birch_pollen"
	GrassPollen                AQIParam = "grass_pollen"
	MugwortPollen              AQIParam = "mugwort_pollen"
	OlivePollen                AQIParam = "olive_pollen"
	RagweedPollen              AQIParam = "ragweed_pollen"
	EuropeanAQI                AQIParam = "european_aqi"
	EuropeanAQIPM25            AQIParam = "european_aqi_pm2_5"
	EuropeanAQIPM10            AQIParam = "european_aqi_pm10"
	EuropeanAQINitrogenDioxide AQIParam = "european_aqi_nitrogen_dioxide"
	EuropeanAQIOzone           AQIParam = "european_aqi_ozone"
	EuropeanAQISulphurDioxide  AQIParam = "european_aqi_sulphur_dioxide"
	USAQI                      AQIParam = "us_aqi"
	USAQIPM25                  AQIParam = "us_aqi_pm2_5"
	USAQIPM10                  AQIParam = "us_aqi_pm10"
	USAQINitrogenDioxide       AQIParam = "us_aqi_nitrogen_dioxide"
	USAQIOzone                 AQIParam = "us_aqi_ozone"
	USAQISulphurDioxide        AQIParam = "us_aqi_sulphur_dioxide"
	USAQICarbonMonoxide        AQIParam = "us_aqi_carbon_monoxide"
)
