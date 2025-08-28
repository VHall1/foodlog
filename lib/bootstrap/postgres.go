package bootstrap

type psqlConfig struct {
	Host         string `envconfig:"POSTGRES_HOST"`
	Username     string `envconfig:"POSTGRES_USERNAME"`
	Password     string `envconfig:"POSTGRES_PASSWORD"`
	DatabaseName string `envconfig:"default=postgres"`
}

// host=postgres port=5432 user=postgres password=postgres dbname=postgres sslmode=disable
