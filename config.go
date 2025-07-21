// Copyright (C) 2025 Özgür Entegrasyon

package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RabbitIP   string `yaml:"rabbit_ip"`
	RabbitPort string `yaml:"rabbit_port"`
	RabbitUser string `yaml:"rabbit_user"`
	RabbitPass string `yaml:"rabbit_pass"`
	Interval   int    `yaml:"interval"`
	Hostname   string `yaml:"hostname"`
}

func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Config read error: %v", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Config parse error: %v", err)
	}
	return cfg
}
