package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UserID  string
	DllPath string
	DirPath string
}

func LoadFromEnv() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return &Config{
		UserID:  os.Getenv("USER_ID"),
		DllPath: os.Getenv("DLL_PATH"),
		DirPath: os.Getenv("DIR_PATH"),
	}
}
