package main

import (
	"fmt"

	"incolore.net/util/str"
)

// GetSensorsByPort 获取某个端口下的所有设备
func GetSensorsByPort(port int) []Sensor {
	var ret []Sensor
	for _, sensor := range BasicInfo.Sensors {
		sensorPort := str.GetPortFromAddr(sensor.LocalAddr)
		//log.Debugf("connect port %v, sensorPort %v", port, sensorPort)

		if port == sensorPort {
			ret = append(ret, sensor)
		}
	}

	return ret
}

// GetSensorByPortDeviceAddr 通过设备地址和端口号获取设备
func GetSensorByPortDeviceAddr(port int, deviceAddr byte) (Sensor, error) {
	for _, sensor := range BasicInfo.Sensors {
		sensorPort := str.GetPortFromAddr(sensor.LocalAddr)
		//log.Debugf("connect port %v, sensorPort %v", port, sensorPort)

		if port == sensorPort && sensor.DeviceAddr == deviceAddr {
			return sensor, nil
		}
	}
	return Sensor{}, fmt.Errorf("Sensor device not found by port %v and addr %x", port, deviceAddr)
}
