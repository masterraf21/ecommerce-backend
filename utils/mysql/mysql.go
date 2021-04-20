package mysql

import (
	"database/sql"

	// _ mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Option can be used to configure mysql connection
type Option struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

// SetupDatabase will prepare mysql connection
func SetupDatabase(readerConfig, writerConfig Option) (reader, writer *sql.DB, err error) {
	reader, err = createConnection(readerConfig)
	if err != nil {
		return nil, nil, err
	}

	writer, err = createConnection(writerConfig)
	if err != nil {
		return nil, nil, err
	}

	return reader, writer, nil
}

func createConnection(config Option) (db *sql.DB, err error) {
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}

	if config.Port == "" {
		config.Port = "3306"
	}

	auth := config.User + ":" + config.Password
	uri := "tcp(" + config.Host + ":" + config.Port + ")"
	dsn := auth + "@" + uri + "/" + config.Database + "?parseTime=true"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(5)

	err = db.Ping()

	return
}
