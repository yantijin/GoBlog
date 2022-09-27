package config

// 数据库相关配置
// 参考文档  https://gorm.io/zh_CN/docs/connecting_to_the_database.html
//username:pwd@tcp(ip:port)/dbname?charset=xxx&parseTime=xxxx&loc=xxx
type CommeDB struct {
	UserName string `json:"username" yaml:"username" mapstructure:"username"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	DdName   string `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
	Port     string `json:"port" yaml:"port" mapstructure:"port"`
	Path     string `json:"path" yaml:"path" mapstructure:"path"`
	Config   string `json:"config" yaml:"config" mapstructure:"config"` // 例如charset=xxx&parseTime=xxxx&loc=xxx
	//下面是连接池的配置
	MaxIdleConns int `json:"max-idle-conns" yaml:"max-idle-conns" mapstructure:"max-idle-conns"` // 空闲连接池中连接的最大数量
	MaxOpenConns int `json:"max-open-conns" yaml:"max-open-conns" mapstructure:"max-open-conns"` // 数据库连接的最大数量
}
