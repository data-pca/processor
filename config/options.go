package config

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"log"
)

var Cfg config

func init() {
	if err := env.Parse(&Cfg); err != nil {
		log.Fatalf("cannot parse env:" + err.Error())
	}
	if err := validator.New().StructCtx(context.Background(), Cfg); err != nil {
		log.Fatalf("cannot validate env:" + err.Error())
	}
}

func (c config) GetMultiplexURL() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func (cfg postgres) GetPgDriverWithConnectionURL() (string, string) {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	return cfg.PGDriver, url
}
