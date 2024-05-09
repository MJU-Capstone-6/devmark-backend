package config

type Config struct {
	DB    DB    `mapstructure:"db"`
	Kakao Kakao `mapstructure:"kakao"`
	App   App   `mapstructure:"app"`
}

type DB struct {
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type App struct {
	Port      string `mapstructure:"port"`
	FooterKey string `mapstructure:"footer_key"`
	IsDevMode bool   `mapstructure:"is_dev_mode"`
}

type Kakao struct {
	ClientKey    string `mapstructure:"client_key"`
	ClientSecret string `mapstructure:"client_secret"`
}
