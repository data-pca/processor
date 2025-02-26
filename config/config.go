package config

type config struct {
	host string `env:"HOST"`
	port uint16 `env:"PORT" validate:"required"`

	Postgres postgres
}

type postgres struct {
	Host     string `env:"HOST" validate:"required"`
	Port     int    `env:"PORT" validate:"required"`
	User     string `env:"USER" validate:"required"`
	Password string `env:"PASSWORD" validate:"required"`
	DBName   string `env:"DB"`
	PGDriver string `env:"DRIVER" validate:"required"`
}
