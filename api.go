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
		Get(Conf.API + "/device_controller/client_config")
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

// ClientInfo ...
type ClientInfo struct {
	UserInfo      UserInfo      `json:"userInfo"`
	SensorConfigs SensorConfigs `json:"sensorConfigs"`
}

// UserInfo ...
type UserInfo struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	ScreenName  string   `json:"screenName"`
	Email       string   `json:"email"`
	Avatar      string   `json:"avatar"`
	Permissions []string `json:"permissions"`
}

// SensorConfig ...
type SensorConfig struct {
	ID           string      `json:"id"`
	UserID       string      `json:"userId"`
	Cycle        string      `json:"cycle"`
	MeterType    string      `json:"meterType"`
	MeasureRange string      `json:"measureRange"`
	LocalAddr    string      `json:"localAddr"`
	DeviceAddr   string      `json:"deviceAddr"`
	Content      string      `json:"content"`
	LimitMax     string      `json:"limitMax"`
	LimitMin     string      `json:"limitMin"`
	Dot          string      `json:"dot"`
	Enabled      string      `json:"enabled"`
	Alias        interface{} `json:"alias"`
	Unit         string      `json:"unit"`
}

// SensorConfigs ...
type SensorConfigs struct {
	Page    int    `json:"page"`
	Perpage int    `json:"perpage"`
	Total   int    `json:"total"`
	Data    []SensorConfig `json:"data"`
}