package config

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