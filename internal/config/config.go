package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Service struct {
	Name           string `yaml:"name"`
	ExpectedStatus string `yaml:"expected_status"`
}

type Notifier struct {
	SMTPServer string   `yaml:"smtpServer"`
	SMTPPort   int      `yaml:"smtpPort"`
	SMTPUser   string   `yaml:"smtp_username"`
	Password   string   `yaml:"password"` //Todo move to env!!
	Recipients []string `yaml:"recipients"`
}

type Config struct {
	Services []Service  `yaml:"services"`
	Notifier []Notifier `yaml:"notifiers"`
	Verbose  bool       `yaml:"verbose"`
	Interval int        `yaml:"interval"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
