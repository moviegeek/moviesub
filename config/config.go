package config

import "github.com/jinzhu/configor"

//Config holds the config values from env variables or config file
type Config struct {
	FileSystem struct {
		MediaRootDir string
	}
	Database struct {
		URI  string
		Auth struct {
			Username string
			Password string
		}
	}
	Server struct {
		Schema     string
		ListenAddr string
		Port       int
	}
}

//Load load the config and return it
func Load() *Config {
	conf := new(Config)
	err := configor.New(&configor.Config{ENVPrefix: "MOVIESUB"}).Load(conf, "config.yaml")
	if err != nil {
		panic(err)
	}

	return conf
}
