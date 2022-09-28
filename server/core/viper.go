package core

import (
	"flag"
	"fmt"
	"os"

	"GoLog/commen"
	"GoLog/core/internal"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 优先级：命令行>环境变量>默认值
func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) > 0 {
		config = path[0]
		fmt.Printf("您正在使用Viper函数传递的参数，config路径为%s", config)
	} else {
		// 先从命令行取
		flag.StringVar(&config, "c", "", "choose config file path.")
		flag.Parse()
		if config == "" {
			// 先从环境变量中获取，如果没有，则从文件中取
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin的%s环境模型，config路径为%s", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin的%s环境模型，config路径为%s", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin的%s环境模型，config路径为%s", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else {
				config = configEnv
				fmt.Printf("您正在使用%s环境变量，config路径为%s", internal.ConfigEnv, config)
			}
		} else {
			fmt.Printf("您正在使用命令行 -c 参数传递的值，config路径为%s", config)
		}
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}
	v.WatchConfig() // 实时监听config文件内容的变化，并及时更新
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件%s发生修改", config)
		if err = v.Unmarshal(&commen.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&commen.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
