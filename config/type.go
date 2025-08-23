package config

type Config struct {
	App    App           `mapstructure:"app" validate:"required"`
	DB     DB            `mapstructure:"db" validate:"required"`
	Redis  Redis         `mapstructure:"redis" validate:"required"`
	Secret SecreteConfig `mapstructure:"secret" validate:"required"`
}

type App struct {
	Port string `mapstructure:"port" validate:"required"`
}

type DB struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Name     string `mapstructure:"name" validate:"required"`
}

type SecreteConfig struct {
	JWTSecret string `mapstructure:"jwt_secret" validate:"required"`
}

type Redis struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	Password string `mapstructure:"password"`
}
