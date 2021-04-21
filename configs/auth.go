package configs

import (
	"os"
)

type auth struct {
	Secret        string
	RefreshSecret string
}

func setupAuth() *auth {
	v := &auth{
		Secret:        os.Getenv("SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}

	return v
}
