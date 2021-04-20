package configs

import (
	"os"
)

type server struct {
	Port string
}

func setupServer() *server {
	v := server{
		Port: os.Getenv("PORT"),
	}

	if v.Port == "" {
		v.Port = "8000"
	}

	return &v
}
