package config

import "strconv"

type System struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

// 获得启动地址
func (s System) GetAddress() string {
	return s.Address + ":" + strconv.Itoa(s.Port)
}
