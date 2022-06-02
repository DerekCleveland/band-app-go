package util

import "github.com/spf13/viper"

type Config struct {
	MongoUsername string `mapstructure:"MONGO_USERNAME"`
	MongoPassword string `mapstructure:"MONGO_PASSWORD"`
	MongoHost     string `mapstructure:"MONGO_HOST"`
	MongoPort     string `mapstructure:"MONGO_PORT"`
	MongoScheme   string `mapstructure:"MONGO_SCHEME"`
}

func LoadConfig() (config Config, err error) {
	viper.BindEnv("MONGO_USERNAME")
	viper.BindEnv("MONGO_PASSWORD")
	viper.BindEnv("MONGO_HOST")
	viper.BindEnv("MONGO_PORT")
	viper.BindEnv("MONGO_SCHEME")

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}
