package config

type JWT struct {
	SigningKey  string `gorm:"signing-key" yaml:"signing-key" mapstructure:"signing-key"`    // jwt签名
	ExpiresTime string `gorm:"expires-time" yaml:"expires-time" mapstructure:"expires-time"` // 过期时间,比如 7d
	Issuer      string `gorm:"issuer" yaml:"issuer" mapstructure:"issuer"`                   // 签发者
}
