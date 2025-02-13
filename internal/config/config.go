package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	InputFilePath    string           `yaml:"input_file_path"`
	OutputFilePath   string           `yaml:"output_file_path"`
	FilterConfig     FilterConfig     `yaml:"filter"`
	FillConfig       FillConfig       `yaml:"fill"`
	AggregatorConfig AggregatorConfig `yaml:"aggregator"`
}

type FilterConfig struct {
	Field     string `yaml:"field"`
	Condition string `yaml:"condition"`
}

type FillConfig struct {
	Field string      `yaml:"field"`
	Value interface{} `yaml:"value"`
}

type AggregatorConfig struct {
	Field        string            `yaml:"field"`
	GroupByField string            `yaml:"group_by_field"`
	Aggregations map[string]string `yaml:"aggregations"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
