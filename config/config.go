package config

import "time"

type config struct {
	host string `env:"HOST"`
	port uint16 `env:"PORT" validate:"required"`

	Postgres postgres `envPrefix:"POSTGRES_"`

	AccessTokenSettings jwt `env:"ACCESS_TOKEN_"`
	RefreshTokenSetting jwt `env:"REFRESH_TOKEN_"`
}

type postgres struct {
	Host     string `env:"HOST" validate:"required"`
	Port     int    `env:"PORT" validate:"required"`
	User     string `env:"USER" validate:"required"`
	Password string `env:"PASSWORD" validate:"required"`
	DBName   string `env:"DB"`
	PGDriver string `env:"DRIVER" validate:"required"`
}

type jwt struct {
	Secret        string        `env:"SECRET"`
	ExpireMinutes time.Duration `env:"EXPIRES_AT"`
}
