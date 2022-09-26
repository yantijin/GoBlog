package config

type Server struct {
	Zap Zap `mapstructure:"level" json:"level" yaml:"string"`
}
