package config

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetConfig(t *testing.T) {
	correct := &Config{
		App{Port: 8080, Api_secret: "secret"},
		Db{Port: 5432, Host: "localhost", User: "postgres", Password: "postgres", DBName: "backend"},
	}
	t.Run("correct config", func(t *testing.T) {
		env := generateConfig(true, true, true, true, true, true, true)
		result, _ := GetConfig(env)
		assert.DeepEqual(t, result, correct)
	})
	t.Run("wrong port", func(t *testing.T) {
		env := generateConfig(false, false, false, false, false, false, false)
		_, err := GetConfig(env)
		assertError(t, err, errApiPort)
	})
	t.Run("wrong api_secret", func(t *testing.T) {
		env := generateConfig(true, false, false, false, false, false, false)
		_, err := GetConfig(env)
		assertError(t, err, errApiSecret)
	})
	t.Run("wrong db_host", func(t *testing.T) {
		env := generateConfig(true, true, false, false, false, false, false)
		_, err := GetConfig(env)
		assertError(t, err, errDbHost)
	})
	t.Run("wrong db_port", func(t *testing.T) {
		env := generateConfig(true, true, true, false, false, false, false)
		_, err := GetConfig(env)
		assertError(t, err, errDbPort)
	})
	t.Run("wrong db_user", func(t *testing.T) {
		env := generateConfig(true, true, true, true, false, false, false)
		_, err := GetConfig(env)
		assertError(t, err, errDbUser)
	})
	t.Run("wrong db_password", func(t *testing.T) {
		env := generateConfig(true, true, true, true, true, false, false)
		_, err := GetConfig(env)
		assertError(t, err, errDbPassword)
	})
	t.Run("wrong db_name", func(t *testing.T) {
		env := generateConfig(true, true, true, true, true, true, false)
		_, err := GetConfig(env)
		assertError(t, err, errDbName)
	})
}

func generateConfig(port, api_secret, db_host, db_port, db_user, db_password, db_name bool) map[string]string {
	env := map[string]string{}
	if port {
		env["port"] = "8080"
	}
	if api_secret {
		env["api_secret"] = "secret"
	}
	if db_host {
		env["db_host"] = "localhost"
	}
	if db_port {
		env["db_port"] = "5432"
	}
	if db_user {
		env["db_user"] = "postgres"
	}
	if db_password {
		env["db_password"] = "postgres"
	}
	if db_name {
		env["db_name"] = "backend"
	}
	return env
}
func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}
