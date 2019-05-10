package config

import "github.com/spf13/viper"

var Current *Config

type Config struct {
	Self       *Server
	BranchAuth struct {
		Host         string
		Port         string
		ClientID     string
		ClientSecret string
		RedirectURI  string
	}
}

type Server struct {
	Host string
	Port string
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Load(fname string) error {
	viper.SetConfigFile(fname)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&Current); err != nil {
		return err
	}

	return nil
}
