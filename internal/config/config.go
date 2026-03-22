package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Site struct {
	URL  string `yaml:"url"`
	Name string `yaml:"name"`
}

type Config struct {
	List    []Site `yaml:"sites"`
	Timeout int    `yaml:"timeout"`
}

func Load(path string) (*Config, error) {
	yfile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer yfile.Close()

	content, err := io.ReadAll(yfile)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
