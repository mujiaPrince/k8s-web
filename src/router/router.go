/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
	"k8s-web/src/api/helloDemo"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	router.Use(cors.New(cfg))
	InitRouter()
}

func InitRouter() {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/hello", helloDemo.Hello)
}

func Start() {
	router.Run(":9999")
}
