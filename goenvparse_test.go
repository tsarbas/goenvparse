package goenvparse

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type config struct {
	Name    string `env:"TEST_NAME" default:"simple"`
	Number  int    `env:"TEST_NUMBER" default:"1"`
	Success bool   `env:"TEST_SUCCESS" default:"true"`
}

func TestParseDefault(t *testing.T) {
	c := &config{}

	err := Parse(c)

	assert.Nil(t, err)

	assert.Equal(t, c.Name, "simple")
	assert.Equal(t, c.Number, 1)
	assert.Equal(t, c.Success, true)
}

func TestParseFile(t *testing.T) {
	c := &config{}

	err := Parse(c, ".test-env")

	assert.Nil(t, err)

	assert.Equal(t, c.Name, "from_file")
	assert.Equal(t, c.Number, 2)
	assert.Equal(t, c.Success, false)
}

func TestParseEnv(t *testing.T) {
	os.Setenv("TEST_NAME", "from_env")
	os.Setenv("TEST_NUMBER", "3")
	os.Setenv("TEST_SUCCESS", "false")

	c := &config{}

	err := Parse(c, ".test-env")

	assert.Nil(t, err)

	assert.Equal(t, c.Name, "from_env")
	assert.Equal(t, c.Number, 3)
	assert.Equal(t, c.Success, false)

	os.Unsetenv("TEST_NAME")
	os.Unsetenv("TEST_NUMBER")
	os.Unsetenv("TEST_SUCCESS")
}

func TestParseNotExistsFile(t *testing.T) {
	c := &config{}
	err := Parse(c, "not-exists")

	assert.Nil(t, err)

	assert.Equal(t, c.Name, "simple")
	assert.Equal(t, c.Number, 1)
	assert.Equal(t, c.Success, true)
}
