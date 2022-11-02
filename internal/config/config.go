package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Port            string `env:"PORT" env-default:":8080"`
	MongoConnection string `env:"MONGO_CONNECTION" env-default:"mongodb://localhost:27017"`
}

func GetConfigs() (*config, error) {
	var cfg config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
