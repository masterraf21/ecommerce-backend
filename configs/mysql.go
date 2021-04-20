package configs

import (
	"fmt"
	"os"
)

type mysql struct {
	Database       string
	ReaderHost     string
	ReaderPort     string
	ReaderUser     string
	ReaderPassword string
	WriterHost     string
	WriterPort     string
	WriterUser     string
	WriterPassword string
}

func setupMySQL() *mysql {
	v := &mysql{}

	v.setupGeneral()
	v.setupReader()
	v.setupWriter()

	return v
}

func (v *mysql) setupGeneral() {
	v.Database = os.Getenv("MYSQL_DATABASE_NAME")

	if v.Database == "" {
		v.Database = "db"
		fmt.Printf("using default %s as MySQL's database name\n", v.Database)
	}
}

func (v *mysql) setupReader() {
	v.ReaderHost = os.Getenv("READER_HOST")
	v.ReaderPort = os.Getenv("READER_PORT")
	v.ReaderUser = os.Getenv("READER_USER")
	v.ReaderPassword = os.Getenv("READER_PASSWORD")

	if v.ReaderHost == "" {
		v.ReaderHost = "127.0.0.1"
		fmt.Printf("using default %s as MySQL's reader host\n", v.ReaderHost)
	}

	if v.ReaderPort == "" {
		v.ReaderPort = "3306"
		fmt.Printf("using default %s as MySQL's reader port\n", v.ReaderPort)
	}

	if v.ReaderUser == "" {
		v.ReaderUser = "reader"
		fmt.Printf("using default %s as MySQL's reader user\n", v.ReaderUser)
	}

	if v.ReaderPassword == "" {
		v.ReaderPassword = "R34D3R_password?"
		fmt.Printf("using default MySQL's reader password\n")
	}
}

func (v *mysql) setupWriter() {
	v.WriterHost = os.Getenv("WRITER_HOST")
	v.WriterPort = os.Getenv("WRITER_PORT")
	v.WriterUser = os.Getenv("WRITER_USER")
	v.WriterPassword = os.Getenv("WRITER_PASSWORD")

	if v.WriterHost == "" {
		v.WriterHost = "127.0.0.1"
		fmt.Printf("using default %s as MySQL's writer host\n", v.WriterHost)
	}

	if v.WriterPort == "" {
		v.WriterPort = "3306"
		fmt.Printf("using default %s as MySQL's writer port\n", v.WriterPort)
	}

	if v.WriterUser == "" {
		v.WriterUser = "writer"
		fmt.Printf("using default %s as MySQL's writer user\n", v.WriterUser)
	}

	if v.WriterPassword == "" {
		v.WriterPassword = "WR1T3R_password?"
		fmt.Printf("using default MySQL's writer password\n")
	}
}
