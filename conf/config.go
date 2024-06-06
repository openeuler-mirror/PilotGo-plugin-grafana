package conf

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gopkg.in/yaml.v3"
)

type HttpConfig struct {
	Addr string `yaml:"addr"`
}

type GrafanaConfig struct {
	Addr string `yaml:"addr"`
}

type ServerConfig struct {
	Http    *HttpConfig     `yaml:"http_server"`
	Grafana *GrafanaConfig  `yaml:"grafana_server"`
	Logopts *logger.LogOpts `yaml:"log"`
}

func config_file() string {
	return "./config.yaml"
}

var global_config ServerConfig

func init() {
	err := readConfig(config_file(), &global_config)
	if err != nil {
		fmt.Printf("")
		os.Exit(-1)
	}
}

func Config() *ServerConfig {
	return &global_config
}

func readConfig(file string, config interface{}) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("open %s failed! err = %s\n", file, err.Error())
		return err
	}

	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		fmt.Printf("yaml Unmarshal %s failed!\n", string(bytes))
		return err
	}
	return nil
}
