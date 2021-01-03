package app

import (
	"errors"
)

// ShowStatusLis ...
func ShowStatusLis(author string) ([]Status, error) {
	var s []Status
	if db2.Debug().Table("status").Where("author = ?", author).Find(&s).RecordNotFound() {
		return nil, errors.New("未找到个人中心的文章")
	}
	return s, nil


}