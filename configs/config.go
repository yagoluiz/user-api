package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port            string `env:"PORT" env-default:":8080"`
	ApiHost         string `env:"API_HOST" env-default:"http://localhost:8080"`
	MongoConnection string `env:"MONGO_CONNECTION" env-default:"mongodb://localhost:27017"`
	Debug           bool   `env:"DEBUG" env-default:"true"`
}

func GetConfigs() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
