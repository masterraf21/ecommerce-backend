package configs

import (
	"fmt"
	"os"
)

type mongodb struct {
	Hosts    string
	Database string
	Options  string
}

func setupMongoDB() *mongodb {
	v := &mongodb{
		Hosts:    os.Getenv("MONGO_HOSTS"),
		Database: os.Getenv("MONGO_DATABASE"),
		Options:  os.Getenv("MONGO_OPTIONS"),
	}

	if v.Database == "" {
		v.Database = "db"
		fmt.Printf("using default %s as MongoDB's database name\n", v.Database)
	}

	return v
}
