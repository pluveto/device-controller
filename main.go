package main

import (
	"fmt"

	log "github.com/cihub/seelog"
	"github.com/go-resty/resty"
)

var Conf *Config
var RestClient *resty.Client
var BasicInfo *ClientInfo

func main() {
	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")
	log.ReplaceLogger(logger)
	// 读取配置文件
	Conf, err = loadConfig("config.json")
	if err != nil {
		log.Error("Failed to load configure file!")
		return
	}
	log.Info(fmt.Sprintf("Config loaded. API Base: %s", Conf.API))
	// 获取传感设备列表
	RestClient = resty.New()
	BasicInfo, err = GetClientInfo()
	if err != nil {
		log.Error("Failed to load client info!")
		return
	}
	log.Info(fmt.Sprintf("Basic info loaded (%v devices). Name: %s", len(BasicInfo.Sensors), BasicInfo.User.Name))
	// 启用定时询查
	go RunService()
	select {}
}
