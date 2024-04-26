package config

type Config struct {
	DB  DB
	App App
}

type DB struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type App struct {
	Port      string `yaml:"port"`
	IsDevMode string `yaml:"is_dev_mode"`
}
