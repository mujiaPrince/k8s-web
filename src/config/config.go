/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package config

import (
	"github.com/spf13/viper"
	"k8s-web/src/go-common/log"
	"path/filepath"
)

// Database ...
type DatabaseConfig struct {
	Driver   string
	Address  string
	DBname   string
	User     string
	Password string
}

/**
function：InitConfig 初始化yaml配置文件
params：
	- string yamlConfigPath 配置文件config.yaml
return:
	- error 返回错误信息
*/
func InitConfig(yamlConfig string) error {
	log.Logger.Infof("InitYamlConfig configpath:%s \n", yamlConfig)
	fullPath, err := filepath.Abs(yamlConfig)
	if err != nil {
		log.Logger.Infof("The file path is error: %s", err.Error())
	}
	log.Logger.Infof("fullPath:%s \n", fullPath)
	viper.SetConfigFile(fullPath)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func GetConfigFile() string {
	return viper.ConfigFileUsed()
}

func GetRestfulListenAddress() string {
	return viper.GetString("server.restful.listenAddress")
}

func GetDBConfig() (db DatabaseConfig, err error) {
	err = viper.UnmarshalKey("Database", &db)
	return
}
