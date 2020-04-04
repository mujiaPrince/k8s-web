/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package context

import (
	"fmt"
	"k8s-web/src/config"
	"k8s-web/src/go-common/log"
	"k8s-web/src/storage"
	"k8s-web/src/storage/db"
)

var (
	server ServerBase
)

type ServerBase interface {
	GetConfigFile() string
	GetDbConnection() storage.DbBackend
}

type Server struct {
	Context *Context
}

// should only be invoke once MUST only in main function!
func InitServer(configFile string, dbConfig *config.DatabaseConfig) error {
	if configFile == "" || dbConfig == nil {
		errInfo := "config file or db config is invalid"
		log.Logger.Infoln(configFile)
		log.Logger.Infoln(dbConfig)
		log.Logger.Infoln(errInfo)
		return fmt.Errorf(errInfo)
	}
	context := &Context{ConfigFile: configFile, DBConfig: dbConfig}
	ser := &Server{Context: context}
	err := ser.initialize()
	server = ser
	return err
}

func InitUintTestServer(serv ServerBase) error {
	server = serv
	return nil
}

// TODO: multi thread race condition refine
func GetServer() ServerBase {
	return server
}

func (s *Server) GetConfigFile() string {
	return s.Context.ConfigFile
}

func (s *Server) GetDbConnection() storage.DbBackend {
	return s.Context.Store
}

func (s *Server) initialize() error {
	db, err := db.NewDbBackend(s.Context.DBConfig)
	if err != nil {
		log.Logger.Infof("new db backend error: %v", err)
		return err
	}
	s.Context.Store = db
	return nil
}
