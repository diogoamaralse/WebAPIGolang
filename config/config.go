package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
}

type Config struct {
	Kafka KafkaConfig `yaml:"kafka"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
