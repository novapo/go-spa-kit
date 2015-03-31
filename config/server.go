package config

// ServerConfig keeps the data for the server configuration
type ServerConfig struct {
	Port int
	Path string
}

var serverConfig *ServerConfig

// GetServerConfig returns a ServerConfig Instance filled with config data
func GetServerConfig() ServerConfig {

	if serverConfig == nil {
		serverConfig = &ServerConfig{8000, "/srv/http"}
	}

	return *serverConfig
}
