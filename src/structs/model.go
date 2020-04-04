/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package structs

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type BlockData struct {
	TxId   string `json:"txId";gorm:"type:varchar(128);primary_key"`
	TxType int32  `json:"txType"`
	TxTime string `json:"txTime"`
}

// Value implements the driver.Valuer interface.
func (bd *BlockData) Value() (driver.Value, error) {
	j, err := json.Marshal(bd)
	return j, err
}

// Scan implements the sql.Scanner interface.
func (bd *BlockData) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	return json.Unmarshal(source, bd)
}

type Student struct {
	ID   int64 `gorm:"type:varchar(128);primary_key"`
	Name string
	Age  int32
	Sex  bool
}
