package settings

import (
	_ "embed"

	// "github.com/go-yaml/yaml"
	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml

var settingFile []byte

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
type Setting struct {
	Port string   `yaml:"port"`
	DB   Database `yaml:"database"`
}

func New() (*Setting, error) {
	var s Setting

	err := yaml.Unmarshal(settingFile, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
