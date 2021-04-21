package configs

// Server config
var Server *server

// MongoDB Config
var MongoDB *mongodb

// Constant config
var Constant *constant

func init() {
	Server = setupServer()
	MongoDB = setupMongoDB()
	Constant = setupConstant()
}
