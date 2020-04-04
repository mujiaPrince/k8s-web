/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	_ "github.com/spf13/viper"
	"k8s-web/src/config"
	"k8s-web/src/context"
	"k8s-web/src/go-common/log"
	"k8s-web/src/router"

	"time"
)

var configPath *string = flag.String("f", "./src/config/config.yaml", "config file name to read!")
var level *string = flag.String("l", "debug", "config log level.")

func main() {
	flag.Parse()

	logLevel, err := logrus.ParseLevel(*level)
	if err != nil {
		log.Logger.Errorf("parse log level error: %v", err)
		return
	}
	err = log.ConfigLogger("./", "log", 30*24*time.Hour,
		10*time.Minute, logLevel, false, true)
	if err != nil {
		log.Logger.Infof("config logger error: %v", err)
		return
	}
	//init mysql
	config.InitConfig(*configPath)

	dbConfig, err := config.GetDBConfig()
	if err != nil {
		log.Logger.Infof("get db config error: %v\n", err)
		return
	}
	err = context.InitServer(config.GetConfigFile(), &dbConfig)
	if err != nil {
		log.Logger.Infof("init server error: %v", err)
		return
	}

	router.Start()

}
