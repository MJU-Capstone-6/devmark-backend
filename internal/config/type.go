package config

type Config struct {
	DB        DB        `mapstructure:"db" json:"db"`
	Kakao     Kakao     `mapsaructure:"kakao" json:"kakao"`
	App       App       `mapstructure:"app" json:"app"`
	OpenAI    OpenAI    `mapstructure:"open_ai" json:"open_ai"`
	GoogleAPI GoogleAPI `mapstructure:"google_api" json:"google_api"`
}

type DB struct {
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type App struct {
	Port             string `mapstructure:"port"`
	FooterKey        string `mapstructure:"footer_key"`
	IsDevMode        bool   `mapstructure:"is_dev_mode"`
	FireBaseFilename string `mapstructure:"firebase_filename"`
}

type Kakao struct {
	ClientKey    string `mapstructure:"client_key"`
	ClientSecret string `mapstructure:"client_secret"`
}

type OpenAI struct {
	ClientKey string `mapstructure:"client_key"`
}

type GoogleAPI struct {
	ClientKey string `mapstructure:"client_key"`
}
