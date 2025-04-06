package config

type Server struct {
	Port string `mapstructure:"SERVER_PORT" validate:"required"`
}
