package config

import (
	"errors"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvData struct  {
	DatabaseURL string
	Port int
}

func LoadEnvData() (EnvData, error) {
	env, err := godotenv.Read();
	
	if (err != nil) {
		return EnvData{}, errors.New("nada")
	}

	url := env["DATABASE_URL"]
	port, err := strconv.Atoi(env["PORT"])

	if err != nil {
		port = 3000
	}

	return EnvData{
		DatabaseURL: url,
		Port: port,
	}, nil
}