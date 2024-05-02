package config

type Config struct {
	DB    DB    `yaml:"db"`
	App   App   `yaml:"app"`
	Kakao Kakao `yaml:"kakao"`
}

type DB struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type App struct {
	Port       string `yaml:"port"`
	IsDevMode  string `yaml:"is_dev_mode"`
	PublicKey  string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
	FooterKey  string `yaml:"footer_key"`
}

type Kakao struct {
	ClientKey    string `yaml:"client_key"`
	ClientSecret string `yaml:"client_secret"`
}
