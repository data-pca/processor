package config

type config struct {
	Host string `env:"HOST"`
	Port uint8  `env:"PORT" validate:"required"`
}
