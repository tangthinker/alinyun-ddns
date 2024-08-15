package config

import "github.com/spf13/viper"

type config struct {
	*viper.Viper
}

var Config = &config{
	Viper: viper.New(),
}

func init() {
	Config.SetConfigName("config")
	Config.SetConfigType("toml")
	Config.AddConfigPath(".")
	Config.AddConfigPath("./config")
	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
