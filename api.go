package main

import (
	"encoding/json"
	"fmt"
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
		//SetAuthToken(Conf.AuthKey).
		Get(Conf.API + "/device_controller/client_config?ak=" + Conf.AuthKey)
	if err != nil {
		//return nil, err
		log.Error("Failed to get devices info!")
		log.Flush()
		os.Exit(1)
	}
	//fmt.Println("Body       :\n", resp)
	log.Infof("Basic Info       :\n%s", resp)
	ret := &ClientInfo{}
	err = json.Unmarshal(resp.Body(), &ret)
	if err != nil {
		log.Error("Failed to parse config:")
		log.Error(resp.Body())
		log.Flush()
		os.Exit(1)
	}
	return ret, nil
}

// Report ...
func Report() {
	var reportData []ReportInfo
	fmt.Print("Report: \n ----------------\n")
	for key, val := range RealtimeData {
		fmt.Printf("%v: %v\n", key, val)
		reportData = append(reportData, ReportInfo{ID: key, Value: val})
	}
	if reportData == nil {
		fmt.Println("skip")
		return
	}
	var wrapper ReportWrapper
	wrapper.Data = reportData
	wrapper.AccessKey = Conf.AuthKey
	body, _ := json.Marshal(wrapper)
	fmt.Printf("reqbody: %s\n", body)
	resp, err := RestClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(Conf.API + "/device_controller/report?ak=" + Conf.AuthKey)
	if err != nil {
		fmt.Printf("Error when reporting: %v", err)
	}
	fmt.Printf("resp: %s", resp.String())
	fmt.Println()
}

// ReportWrapper ...
type ReportWrapper struct {
	Data      []ReportInfo `json:"data"`
	AccessKey string       `json:"ak"`
}

// ReportInfo ...
type ReportInfo struct {
	ID    int `json:"id"`
	Value int `json:"value"`
}

// ClientInfo ...
type ClientInfo struct {
	UserInfo      UserInfo       `json:"userInfo"`
	SensorConfigs []SensorConfig `json:"sensorConfigs"`
}

// UserInfo ...
type UserInfo struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	ScreenName  string   `json:"screenName"`
	Email       string   `json:"email"`
	Avatar      string   `json:"avatar"`
	Permissions []string `json:"permissions"`
	Cycle       int      `json:"cycle"`
}

// SensorConfig ...
type SensorConfig struct {
	ID           int         `json:"id,string"`
	UserID       int         `json:"userId,string"`
	Cycle        int         `json:"cycle,string"`
	MeterType    int         `json:"meterType,string"`
	MeasureRange string      `json:"measureRange"`
	LocalAddr    string      `json:"localAddr"`
	DeviceAddr   int         `json:"deviceAddr,string"`
	Content      string      `json:"content"`
	LimitMax     int         `json:"limitMax,string"`
	LimitMin     int         `json:"limitMin,string"`
	Dot          int         `json:"dot,string"`
	Enabled      int         `json:"enabled,string"`
	Alias        interface{} `json:"alias"`
	Unit         string      `json:"unit"`
}
