/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"k8s-web/src/config"
	"k8s-web/src/structs"
	"testing"
)

func TestSetStudent(t *testing.T) {

	stu := &structs.Student{
		Name: "zy",
		Age:  18,
		Sex:  true,
	}
	dbConfig, _ := config.GetDBConfig()
	db, err := NewDbBackend(&dbConfig)
	if err != nil {
		t.Errorf("NewDbBackend error: %v", err)
	}
	db.SetStudent(stu)

}

func TestNewDbBackend(t *testing.T) {
	// Database ...
	type DatabaseConfig struct {
		Driver   string
		Address  string
		DBname   string
		User     string
		Password string
	}
	dbConfig := DatabaseConfig{
		Driver:   "mysql",
		Address:  "114.67.79.102:3306",
		DBname:   "go_kass",
		User:     "root",
		Password: "12345",
	}

	connUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Address,
		dbConfig.DBname,
	)
	// connUrl = "root:zaq1xsw@@tcp(10.0.90.153:33002)/yhbaas?charset=utf8&parseTime=true&loc=Local"
	fmt.Println(connUrl)
	db, err := gorm.Open("mysql", connUrl)
	if err != nil {
		fmt.Println(err)
		panic("连接数据库失败")
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&structs.Student{})

	// 创建
	db.Create(&structs.Student{Name: "L1212", Age: 1000, Sex: false})

	// 读取
	var product structs.Student
	db.First(&product, 1)                   // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Name", "L1212")

	// 删除 - 删除product
	db.Delete(&product)
}
