package config

// 后端配置入口，可以通过Viper读取对应配置文件，对其中所有的属性进行初始化
type Server struct {
	Zap Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Db  CommeDB `mapstructure:"db" json:"db" yaml:"db"`
	JWT JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
