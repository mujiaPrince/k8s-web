/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"k8s-web/src/config"
	"k8s-web/src/go-common/log"
	"k8s-web/src/structs"
)

type DbBackend struct {
	Db *gorm.DB
}

func NewDbBackend(dbConf *config.DatabaseConfig) (*DbBackend, error) {
	connUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Address,
		dbConf.DBname,
	)
	db, err := gorm.Open(dbConf.Driver, connUrl)
	//	db.LogMode(true)

	if err != nil {
		log.Logger.Infof("open db error: %v\n", err)
		return nil, err
	}
	// TODO: this cost too much time,will be refined soon
	db.AutoMigrate(&structs.Student{})
	dbBackend := &DbBackend{Db: db}
	return dbBackend, nil
}

func (db *DbBackend) SetStudent(student *structs.Student) error {
	_, err := db.GetStudent(student.ID, student.Name)
	if err == nil {
		blockInfoModel := &structs.Student{ID: student.ID, Name: student.Name}
		return db.Db.Model(blockInfoModel).UpdateColumn(student).Error
	}
	return db.Db.Create(student).Error
}

func (db *DbBackend) GetStudent(studentId int64, name string) (*structs.Student, error) {
	blockInfo := &structs.Student{}
	err := db.Db.Find(blockInfo, structs.Student{ID: studentId, Name: name}).Error
	if err != nil {
		return nil, err
	}
	return blockInfo, nil
}
