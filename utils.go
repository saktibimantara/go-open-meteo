package go_open_meteo

func safeGetFloat64(slice []float64, index int) *float64 {
	if index >= 0 && index < len(slice) {
		return &slice[index]
	}
	return nil
}

func safeGetWeatherCode(slice []WeatherCodeResponse, index int) *WeatherCodeResponse {
	if index >= 0 && index < len(slice) {
		return &slice[index]
	}
	return nil
}

func safeGetInt(slice []int, index int) *int {
	if index >= 0 && index < len(slice) {
		return &slice[index]
	}
	return nil
}

func safeGetString(slice []string, index int) *string {
	if index >= 0 && index < len(slice) {
		return &slice[index]
	}
	return nil
}
