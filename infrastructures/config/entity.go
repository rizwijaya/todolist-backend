package config

type LoadConfig struct {
	App struct {
		Mode       string `env:"APP_MODE"`
		Name       string `env:"APP_NAME"`
		Port       string `env:"APP_PORT"`
		Url        string `env:"APP_URL"`
		Secret_key string `env:"APP_SECRET_KEY"`
	}

	Database struct {
		Host     string `env:"MYSQL_HOST"`
		Port     string `env:"MYSQL_PORT"`
		User     string `env:"MYSQL_USER"`
		Password string `env:"MYSQL_PASSWORD"`
		Dbname   string `env:"MYSQL_DBNAME"`
	}
}
