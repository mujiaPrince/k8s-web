/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */

package helloDemo

import (
	"github.com/gin-gonic/gin"
	"k8s-web/src/go-common/log"
)

//Hello ...
func Hello(c *gin.Context) {
	log.Logger.Println("hello world")
}
