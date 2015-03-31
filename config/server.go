package config

// ServerConfig keeps the data for the server configuration
type ServerConfig struct {
	Port int
	Path string
}

// GetServerConfig returns a ServerConfig Instance filled with config data
func GetServerConfig() ServerConfig {
	return ServerConfig{8000, "/var/www"}
}
