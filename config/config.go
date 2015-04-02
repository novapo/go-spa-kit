package config

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"
)

// AdminConfig holds config data for the backend
type AdminConfig struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// ServerConfig holds config data for the server configuration
type ServerConfig struct {
	Port int    `json:"port"`
	Path string `json:"path"`
}

// PlugInConfig interface undefined
type PlugInConfig interface{}

// Config holds config-data for server, admin and plugin
type Config struct {
	Admin   AdminConfig  `json:"admin"`
	Server  ServerConfig `json:"server"`
	PlugIns PlugInConfig `json:"plugins"`
}

// AdnConf holds admin related data
var AdnConf *AdminConfig

// ServerConf holds server related data
var ServerConf *ServerConfig

// PlugInConf holds plugin related data
var PlugInConf *PlugInConfig

// Conf holds global config data
var Conf *Config

func init() {
	Conf, _ = ReadConfig()

	// Programm ausführen sinnvoll?
	if Conf == nil {
		syscall.Exit(1)
	}

	AdnConf = &Conf.Admin
	ServerConf = &Conf.Server
	PlugInConf = &Conf.PlugIns

}

// ReadConfig reads data of admin, server and plug-ins from backend
func ReadConfig() (*Config, error) {

	file, err := os.Open("config/config.json")
	if err != nil {
		fmt.Errorf("unable to open config.json/n/tErrorcode: (%v)", err)
		return nil, err
	}

	config := &Config{}

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(config); err != nil {
		fmt.Errorf("Malformated json/n/tErrorcode: (%v)", err)
		return nil, err
	}

	return config, nil
}

// GetAdminConfig returns admin related data
func GetAdminConfig() *AdminConfig {
	return AdnConf
}

// GetServerConfig returns server related data
func GetServerConfig() *ServerConfig {
	return ServerConf
}

// GetPlugInConfig returns plugin related data
func GetPlugInConfig() *PlugInConfig {
	return PlugInConf
}
