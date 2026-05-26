package config

import (
	"os"
	"strconv"

	"gopkg.in/ini.v1"
)

type Config struct {
	HTTPHost         string
	HTTPPort         string
	PostgresHost     string
	PostgresPort     uint16
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
}

func Load(path string) (*Config, error) {
	cfg := &Config{}

	file, err := ini.LooseLoad(path)
	if err != nil {
		return nil, err
	}

	section := file.Section("http")
	cfg.HTTPHost = getKey(section, "host", "JAM_HTTP_HOST", "0.0.0.0")
	cfg.HTTPPort = getKey(section, "port", "JAM_HTTP_PORT", "8080")

	section = file.Section("postgres")
	cfg.PostgresHost = getKey(section, "host", "JAM_POSTGRES_HOST", "localhost")
	port, err := strconv.Atoi(getKey(section, "port", "JAM_POSTGRES_PORT", "5432"))
	if err != nil {
		return nil, err
	}
	cfg.PostgresPort = uint16(port)
	cfg.PostgresDB = getKey(section, "db", "JAM_POSTGRES_DB", "jam")
	cfg.PostgresUser = getKey(section, "user", "JAM_POSTGRES_USER", "jam")
	cfg.PostgresPassword = getKey(section, "password", "JAM_POSTGRES_PASSWORD", "jam")

	if err := file.SaveTo(path); err != nil {
		return nil, err
	}

	return cfg, nil
}

func getKey(section *ini.Section, name string, env_name string, fallback string) string {
	if env := os.Getenv(env_name); env != "" {
		return env
	}

	key := section.Key(name)

	if key.String() == "" {
		key.SetValue(fallback)
		return fallback
	}

	return key.Value()
}
