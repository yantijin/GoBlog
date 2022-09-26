package config

type Zap struct {
	Level        string `json:"level" yaml:"level" mapstructure:"level"`                          // 日志级别
	EncodeLevel  string `json:"encodeLevel" yaml:"encode-level" mapstructure:"encode-level"`      // 编码级别
	Format       string `json:"format" yaml:"format" mapstructure:"format"`                       //encode输出的类型是json格式还是plain text格式
	LogInConsole bool   `json:"logInConsole" yaml:"log-in-console" mapstructure:"log-in-console"` // 是否将日志输出到控制台
	Directory    string `json:"directory" yaml:"directory" mapstructure:"directory"`              // 日志存储的文件夹名称
}
