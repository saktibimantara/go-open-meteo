package go_open_meteo

import (
	"encoding/json"
	"fmt"
	go_http "github.com/saktibimantara/go-http"
)

type GoOpenMeteo struct {
	config  IConfig
	callApi go_http.CallAPI
}

func New(config *Config) *GoOpenMeteo {
	caller := go_http.New(&go_http.Config{})
	return &GoOpenMeteo{config: config, callApi: caller}
}

type IGoOpenMeteo interface {
	Forecast(param IForecastParams) (*ForecastResponse, error)
}

func (g *GoOpenMeteo) Forecast(param IForecastParams) (*ForecastResponse, error) {
	url := g.config.GetForecastURL() + "?" + param.GetParams()
	fmt.Println(url)
	callResp, err := g.callApi.Get(url)
	if err != nil {
		return nil, err
	}

	if callResp.Code != 200 {
		return nil, fmt.Errorf("error code: %d", callResp.Code)
	}

	var resp ForecastResponse

	err = json.Unmarshal(callResp.Data, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
