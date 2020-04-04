/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package context

import (
	"k8s-web/src/config"
	"k8s-web/src/storage"
)

var (
	serverContext *Context
)

type Context struct {
	Store      storage.DbBackend
	ConfigFile string
	DBConfig   *config.DatabaseConfig
}
