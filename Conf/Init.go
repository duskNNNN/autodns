package confGet

import (
	"github.com/spf13/viper"
	"os"
)

// 配置文件设置
func ConfigInit() {
	// 获取当前的工作路径
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// 设置配置文件的名称
	viper.SetConfigName("conf")
	// 设置配置文件的类型
	viper.SetConfigType("yml")
	// 设置配置文件的路径
	viper.AddConfigPath(workDir + "/Conf")
	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
