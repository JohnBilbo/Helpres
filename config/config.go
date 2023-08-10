package config

type Config struct {
	HTTPServer HTTPServer
	Database   Database
}

func ReadConfigFromEnv() Config {
	viper.AutomaticEnv()
	var config Config

	// Server
	config.HTTPServer.HOST = viper.GetString("APP.HTTP_SERVER.HOST")
	config.HTTPServer.PORT = viper.GetString("APP.HTTP_SERVER.PORT")

	// Database
	config.HTTPServer.PORT = viper.GetString("APP.DATABASE.HOST")
	config.HTTPServer.PORT = viper.GetString("APP.DATABASE.PORT")
	config.HTTPServer.PORT = viper.GetString("APP.DATABASE.DB_NAME")
	config.HTTPServer.PORT = viper.GetString("APP.DATABASE.USER")
	config.HTTPServer.PORT = viper.GetString("APP.DATABASE.PASSWORD")
	config.HTTPServer.PORT = viper.GetString("APP.DATABASE.SSL_MODE")
	config.HTTPServer.PORT = viper.GetString("APP.DATABASE.TIME_ZONE")

	return config
}

// TODO - add ReadConfigFromFile
