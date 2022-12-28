package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

type configToml struct {
	Viewer Viewer
	System SystemConfig
}

var Conf *configToml

func init() {
	Conf = new(configToml)
	_, err := toml.DecodeFile("config/config.toml", &Conf)
	if err != nil {
		log.Println(err)
	}
	Conf.System.AppName = "appName"
	Conf.System.Version = 1.0
	path, err := os.Getwd()
	Conf.System.CurrentDir = path
	if err != nil {
		log.Println(err)
	}
}
