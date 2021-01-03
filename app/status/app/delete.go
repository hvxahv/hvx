package app

import (
	"errors"
)


func DeleteStatusByID(id string) error {
	if err := db2.Debug().Table("status").Where("id = ?", id).Delete(&Status{}).Error; err != nil {
		return errors.New("Delete Status Error")
	}

	return nil
}