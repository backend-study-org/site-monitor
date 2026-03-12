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

	bytes, err2 := io.ReadAll(yfile)
	if err2 != nil {
		return nil, err2
	}

	var data Config
	err3 := yaml.Unmarshal(bytes, &data)

	if err3 != nil {
		return nil, err
	}

	return &data, nil
}
