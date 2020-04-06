# goenvparse
Golang structure parser from environment variables and .env files with default values


## Usage

    package main

    import "github.com/tsarbas/goenvparse"

    type Config struct {
        Name    string `env:"TEST_NAME" default:"simple"`
        Number  int    `env:"TEST_NUMBER" default:"1"`
        Success bool   `env:"TEST_SUCCESS" default:"true"`
    }

    c := &Config{}
    goenvparse.Parse(c, ".env", "/path/to/file")
