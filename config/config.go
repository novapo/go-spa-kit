package config

var AdnConf *AdminConfig

func init() {
	AdnConf, _ = ReadAdminConf()
}
