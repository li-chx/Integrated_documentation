package model

import "gorm.io/gorm"

type CarShare struct {
	Userid      int    `json:"userid"gorm:"userid"`
	Id          int    `json:"id"gorm:"id;primary_key;auto_increment"`
	Require     string `json:"require"`
	Destination string `json:"destination"`
	Address     string `json:"address"`
	Begintime   string `json:"begintime"`
	Endtime     string `json:"endtime"`
	Pending     []int  `json:"pending,omitempty"`
	Member      []int  `json:"member,omitempty"`
	Num         int    `json:"num"`
	Mannum      int    `json:"mannum"`
	Womannum    int    `json:"womannum"`
	Maxnum      int    `json:"maxnum"`
	Status      string `json:"status"`
}

func CreateCarShare(sharer *CarShare) error {
	return db.Create(sharer).Error
}

func GetCarShareByDestination(destination string) (CarShare, error) {
	var sharer CarShare
	err := db.Where("destination=?", destination).First(&sharer).Error
	return sharer, err
}

func GetCarShareById(CarShareId int) (CarShare, error) {
	var sharer CarShare
	err := db.Where("id=?", CarShareId).First(&sharer).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return sharer, err
}

func UpdateCarShare(sharer CarShare, mp interface{}) error {
	return db.Model(&sharer).Updates(mp).Error
}

func GetAllCarShare() ([]CarShare, int, error) {
	sharers := make([]CarShare, 1000) //CarShare模型切片
	var c int64
	err := db.Find(&sharers).Count(&c).Error
	//Count函数，直接返回查询匹配的行数。
	return sharers, int(c), err
}

func DeleteCarShare(sharer CarShare) error {
	//删除指定用户以及数据库中的删除记录
	err := db.Unscoped().Delete(&sharer).Error
	return err
}
