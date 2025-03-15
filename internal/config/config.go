package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type TgCfg struct{
	Use bool `yaml:"use"`
	ChatId int64 `yaml:"chat_id"`
	Token string `yaml:"token"`
}

type EmailCfg struct{
	Use bool `yaml:"use"`
	From string `yaml:"from"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
}

type SmsCfg struct{
	Use bool `yaml:"use"`
	From string `yaml:"from"`
}

type Config struct{
	KafkaUrl string `yaml:"kafka_url"`
	EmailCfg EmailCfg `yaml:"email"`
	SmsCfg SmsCfg `yaml:"sms"`
	TgCfg TgCfg `yaml:"tg"`
}

func Init() (*Config,error) {
	file,err := os.Open("config.yaml")
	if err != nil{
		return nil, fmt.Errorf("cant open config file: %v",err)
	}

	body,err := io.ReadAll(file)
	if err != nil{
		return nil, fmt.Errorf("cant get body: %v", err)
	}

	var cfg Config

	if err = yaml.Unmarshal(body, &cfg);err != nil{
		return nil, fmt.Errorf("cant unmarshal: %v",err)
	}

	return &cfg, nil
}