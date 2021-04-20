package configs

// Server config
var Server *server

// MongoDB Config
var MongoDB *mongodb

func init() {
	Server = setupServer()
	MongoDB = setupMongoDB()
}
