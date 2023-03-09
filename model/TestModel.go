package model

import "gorm.io/gorm"

type Test struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}

func Creat(test *Test) error {
	//在数据库中创建一个新用户
	return db.Create(test).Error
}

func GetById(Id int) (Test, error) {
	var test Test
	err := db.Where("id=?", Id).First(&test).Error //找到用户并赋值
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return test, err
}
