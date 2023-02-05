package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig) //用于保存全局配置信息

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
	Port     int    `mapstructure:"port"`
	MaxCons  int    `mapstructure:"max_cons"`
	MaxIdles int    `mapstructure:"max_idles"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Dd       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigName("config")
	//viper.SetConfigType("yaml")  //专用于读取远程配置
	viper.AddConfigPath(".")

	// 读取配置信息 (读到 Conf 中(仅读一次))
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return err
	}

	// 反系列化配置
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarchel failed, err:%v\n", err)
	}

	//热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("The config was been modified.")

		// Dolist 添加自定义通知

		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshel failed, err:%v\n", err)
		}
	})
	return
}
