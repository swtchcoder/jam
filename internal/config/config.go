package config

import (
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	Host string
	Port string
}

func Load(path string) (*Config, error) {
	cfg := &Config{}

	file, err := ini.LooseLoad(path)
	if err != nil {
		return nil, err
	}

	section := file.Section("http")
	cfg.Host = getKey(section, "host", "JAM_HTTP_HOST", "0.0.0.0")
	cfg.Port = getKey(section, "port", "JAM_HTTP_PORT", "8080")

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
