package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"udns/logger"
)

const (
	defaultCfgFile = "udns.cfg.yml"
)

var (
	Cfg *config
)

type config struct {
	Proto     string   `yaml:"proto"`
	Port      int      `yaml:"port"`
	LogLevel  string   `yaml:"logLevel"`
	TTL       int      `yaml:"ttl"`
	Zones     []string `yaml:"zones"`
	ParentDNS string   `yaml:"parent"`
	MyIP      string   `yaml:"myip"`
	Source    string   `yaml:"source"`
}

func Init(cfgFile string) {
	if cfgFile == "" {
		logger.Infof("config", "config file is not set, use %s", defaultCfgFile)
		cfgFile = defaultCfgFile
	}
	data, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		logger.Fatal("config", err)
	}
	if err = yaml.Unmarshal(data, Cfg); err != nil {
		logger.Fatal("config", err)
	}
}
