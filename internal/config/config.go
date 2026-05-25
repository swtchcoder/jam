package config

import (
	"gopkg.in/ini.v1"
	"os"
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

	section := file.Section("")
	cfg.Host = getKey(section, "host", "0.0.0.0")
	cfg.Port = getKey(section, "port", "8080")

	if err := file.SaveTo(path); err != nil {
		return nil, err
	}

	return cfg, nil
}

func getKey(section *ini.Section, name string, fallback string) string {
	if env := os.Getenv(name); env != "" {
		return env
	}

	key := section.Key(name)

	if key.String() == "" {
		key.SetValue(fallback)
		return fallback
	}

	return key.Value()
}
