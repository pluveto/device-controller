package main

import (
	"fmt"

	"incolore.net/util/str"
)

// GetSensorsByPort 获取某个端口下的所有设备
func GetSensorsByPort(port int) []SensorConfig {
	var ret []SensorConfig
	for _, sensorConfig := range BasicInfo.SensorConfigs {
		sensorPort := str.GetPortFromAddr(sensorConfig.LocalAddr)
		//log.Debugf("connect port %v, sensorPort %v", port, sensorPort)

		if port == sensorPort {
			ret = append(ret, sensorConfig)
		}
	}

	return ret
}

// GetSensorByPortDeviceAddr 通过设备地址和端口号获取设备
func GetSensorByPortDeviceAddr(port int, deviceAddr byte) (SensorConfig, error) {
	for _, sensorConfig := range BasicInfo.SensorConfigs {
		sensorPort := str.GetPortFromAddr(sensorConfig.LocalAddr)
		//log.Debugf("connect port %v, sensorPort %v", port, sensorPort)
		if port == sensorPort && byte(sensorConfig.DeviceAddr) == deviceAddr {
			return sensorConfig, nil
		}
	}
	return SensorConfig{},
		fmt.Errorf("Sensor device not found by port %v and addr %x",
			port, deviceAddr)
}
