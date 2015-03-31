package config

// AdnConf holds admin related data
var AdnConf *AdminConfig

func init() {
	AdnConf, _ = ReadAdminConf()
}
