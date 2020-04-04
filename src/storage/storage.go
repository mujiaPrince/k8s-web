/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package storage

import (
	"k8s-web/src/structs"
)

type DbBackend interface {
	SetStudent(student *structs.Student) error
	GetStudent(studentId int64, name string) (*structs.Student, error)
}
