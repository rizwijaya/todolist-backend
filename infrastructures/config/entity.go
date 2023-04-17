package config

type App struct {
	Mode string `env:"APP_MODE"`
	Name string `env:"APP_NAME"`
	Port string `env:"APP_PORT"`
	Url  string `env:"APP_URL"`
}

type Database struct {
	Host     string `env:"MYSQL_HOST"`
	Port     string `env:"MYSQL_PORT"`
	User     string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
	Dbname   string `env:"MYSQL_DBNAME"`
}
type LoadConfig struct {
	App      App
	Database Database
}
