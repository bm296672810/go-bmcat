package config

import (
	lg "bmcat/src/daemon/log"
	"encoding/json"
	"os"
)

type SqliteConf struct {
	Path string `json:"path"`
}

type HttpServerConf struct {
	ListenPort string `json:"listen_port"`
}

type Conf struct {
	SqliteConf     `json:"sqlite"`
	HttpServerConf `json:"httpserver"`
}

func (c *Conf) InitConf() {
	c.Path = "./database.db"
	c.ListenPort = ":2200"
}

func LoadConfig(configPath string) {
	conf, err := os.ReadFile(configPath)
	if err != nil {
		lg.ELogger.Printf("ReadConfigure error:%v\n", err)
		GloabalConf.InitConf()
		return
	}
	// config.ILogger.Println(string(conf))
	err = json.Unmarshal(conf, &GloabalConf)
	if err != nil {
		lg.ELogger.Printf("unmarshal error:%v\n", err)
		// config.GloabalApiConf.InitConfig()
		GloabalConf.InitConf()
	}
}

var GloabalConf Conf
