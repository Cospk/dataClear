package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var (
	cfg *Config
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

func LoadConfigByViper() (*Config, error) {

	// 使用viper加载配置信息
	config := viper.New()
	config.AddConfigPath("/internal/config")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已修改:", e.Name)
		if err := config.Unmarshal(&cfg); err != nil {
			fmt.Println(err)
		}
	})

	// 将配置文件内容解析到config结构体中
	if err := config.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
	}

	return cfg, nil
}
