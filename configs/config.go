package configs

// Server config
var Server *server

// MySQL config
var MySQL *mysql

func init() {
	MySQL = setupMySQL()
	Server = setupServer()
}
