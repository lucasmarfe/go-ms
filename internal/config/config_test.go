package config_test

import (
	"github.com/stretchr/testify/assert"
	"go-proj/internal/config"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tempDir := t.TempDir()
	configYML := []byte(`
server:
  port: 8080
database:
  host: "localhost"
  port: 5432
  user: "testuser"
  password: "testpass"
  dbname: "testdb"
  sslmode: "disable"
`)

	err := os.WriteFile(filepath.Join(tempDir, "config.yml"), configYML, 0644)
	if err != nil {
		t.Fatalf("Failed to write test config file: %v", err)
	}

	cfg := config.LoadConfigFromPath(tempDir)

	// Assert values
	assert.Equal(t, 8080, cfg.Server.Port)
	assert.Equal(t, "localhost", cfg.Database.Host)
	assert.Equal(t, 5432, cfg.Database.Port)
	assert.Equal(t, "testuser", cfg.Database.User)
	assert.Equal(t, "testpass", cfg.Database.Password)
	assert.Equal(t, "testdb", cfg.Database.DBName)
	assert.Equal(t, "disable", cfg.Database.SSLMode)
}
