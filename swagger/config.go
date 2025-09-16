package swagger

import "github.com/Phofuture/photon-core-starter/configuration"

func init() {
	configuration.Register(&config)
}

type Config struct {
	Swagger struct {
		DirectoryPath string `mapstructure:"directory-path"`
		ContextPath   string `mapstructure:"context-path"`
		Enabled       bool   `mapstructure:"enabled"`
	} `mapstructure:"swagger"`
}

var config Config
