package main

import (
	"encoding/json"
	"os"

	log "github.com/cihub/seelog"
)

// GetClientInfo 获取客户端信息
func GetClientInfo() (*ClientInfo, error) {
	if nil == Conf {
		log.Error("Config is not loaded!")
		log.Flush()
		os.Exit(1)
	}
	if nil == RestClient {
		log.Error("Restful client is not prepared!")
		log.Flush()
		os.Exit(1)
	}

	resp, err := RestClient.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetAuthToken(Conf.AuthKey).
		Get(Conf.API + "/config")
	if err != nil {
		//return nil, err
		log.Error("Failed to get devices info!")
		os.Exit(1)
	}
	//fmt.Println("Body       :\n", resp)
	log.Infof("Basic Info       :\n%s", resp)
	ret := &ClientInfo{}
	json.Unmarshal(resp.Body(), &ret)
	return ret, nil
}

// ClientInfo 客户端信息
type ClientInfo struct {
	User struct {
		ID       int    `json:"id"`
		Interval int    `json:"interval"`
		Name     string `json:"name"`
	} `json:"user"`
	Sensors []Sensor `json:"sensors"`
}

// Sensor 传感器配置
type Sensor struct {
	ID           int     `json:"id"`
	MeterType    string  `json:"meterType"`
	Enabled      bool    `json:"enabled"`
	DeviceAddr   uint8   `json:"deviceAddr"`
	Gas          string  `json:"gas"`
	Max          float64 `json:"max"`
	Min          float64 `json:"min"`
	MessureRange string  `json:"messureRange"`
	LocalAddr    string  `json:"localAddr"`
	Building     string  `json:"building"`
	Floor        string  `json:"floor"`
	Alias        string  `json:"alias"`
}
