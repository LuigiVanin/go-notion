package config

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvData struct {
	DatabaseURL string
	JwtKey      string
	Port        int
	Salt        int
}

var GlobalEnv *EnvData

func LoadEnvData() (EnvData, error) {
	env, err := godotenv.Read()

	if err != nil {
		return EnvData{}, errors.New("nada")
	}

	url := env["DATABASE_URL"]

	fmt.Println("ENV HWT_KEY: ", env["JWT_KEY"])
	jwtKey := env["JWT_KEY"]
	salt, err := strconv.Atoi(env["SALT"])

	if err != nil {
		salt = 10
	}

	port, err := strconv.Atoi(env["PORT"])

	if err != nil {
		port = 3000
	}

	GlobalEnv = &EnvData{
		DatabaseURL: url,
		Port:        port,
		Salt:        salt,
		JwtKey:      jwtKey,
	}

	return *GlobalEnv, nil
}
