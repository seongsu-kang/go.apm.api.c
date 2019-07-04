package model

import (
	"fmt"
	"sync"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Service   string
	Logpaths  logpaths
	Logconfig logconfig
	Databases map[string]databases
	Servers   map[string]servers
}

type logpaths struct {
	Logpath string
}

type logconfig struct {
	Logpath  string
	Loglevel int
}

type servers struct {
	IP   string
	PORT string
}

type databases struct {
	Server string
	Port   string
	Enable bool
}



var config TomlConfig

func  Load (cp string) {
	cmdargs := GetCmdargs()
	fpath := fmt.Sprintf(cp, cmdargs.Phase)
	if _, err := toml.DecodeFile(fpath, &config); err != nil {
		fmt.Println(err)
	}
}

func (t *TomlConfig) ApmServerUrl() string {
	return fmt.Sprintf("%s%s", t.Servers["APM"].IP, t.Servers["APM"].PORT)
}

var insTomlConfig *TomlConfig
var onceTomlConfig sync.Once

func GetConfig() *TomlConfig {
	onceTomlConfig.Do(func() {
		insTomlConfig = &config
	})
	return insTomlConfig
}