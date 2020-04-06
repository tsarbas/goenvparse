package goenvparse

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/mcuadros/go-defaults"
)

func Parse(cfg interface{}, filenames ...string) error {
	defaults.SetDefaults(cfg)

	if len(filenames) > 0 {
		if err := godotenv.Load(filenames...); err != nil {
			return err
		}
	}

	if err := env.Parse(cfg); err != nil {
		return err
	}

	return nil
}
