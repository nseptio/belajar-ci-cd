package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig_Valid(t *testing.T) {

	err := os.Setenv("SERVER_PORT", "8080")
	assert.NoError(t, err)
	err = os.Setenv("DB_HOST", "localhost")
	assert.NoError(t, err)
	err = os.Setenv("DB_PORT", "5432")
	assert.NoError(t, err)
	err = os.Setenv("DB_USERNAME", "postgres")
	assert.NoError(t, err)
	err = os.Setenv("DB_PASSWORD", "password")
	assert.NoError(t, err)
	err = os.Setenv("DB_NAME", "testDB")
	assert.NoError(t, err)

	cfg := NewConfig()
	assert.Equal(t, 8080, cfg.Server.Port)
	assert.Equal(t, "localhost", cfg.DB.Host)
	assert.Equal(t, 5432, cfg.DB.Port)
	assert.Equal(t, "postgres", cfg.DB.Username)
	assert.Equal(t, "password", cfg.DB.Password)
	assert.Equal(t, "testDB", cfg.DB.DBName)
}

func TestNewConfig_MissingEnv(t *testing.T) {
	err := os.Unsetenv("SERVER_PORT")
	assert.NoError(t, err)
	err = os.Unsetenv("DB_HOST")
	assert.NoError(t, err)
	err = os.Unsetenv("DB_PORT")
	assert.NoError(t, err)
	err = os.Unsetenv("DB_USERNAME")
	assert.NoError(t, err)
	err = os.Unsetenv("DB_PASSWORD")
	assert.NoError(t, err)
	err = os.Unsetenv("DB_NAME")
	assert.NoError(t, err)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when SERVER_PORT was missing")
		}
	}()

	// Expect the application to panic due to missing required environment variables
	_ = NewConfig()
}
