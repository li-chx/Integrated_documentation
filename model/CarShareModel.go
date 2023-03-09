package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type CarShare struct {
	Userid      int                    `json:"userid" gorm:"userid"`
	Id          int                    `json:"id" gorm:"id;primary_key;auto_increment"`
	Begintime   string                 `json:"begintime"`
	Address     string                 `json:"address"`
	Destination string                 `json:"destination"`
	Num         int                    `json:"num"`
	Maxnum      int                    `json:"maxnum"`
	Luggage     string                 `json:"luggage"`
	Luggages    map[string]interface{} `json:"luggages" gorm:"-"`
	Lugg        string                 `json:"-" gorm:"column:luggages"`
	Box         int                    `json:"box"`
	Bag         int                    `json:"bag"`
	Contact     string                 `json:"contact"`
}

func (data *CarShare) BeforeSave() error {
	if data.Luggages == nil {
		return nil
	}

	b, err := json.Marshal(&data.Luggages)
	if err != nil {
		return err
	}

	data.Lugg = string(b)
	return nil
}

// AfterFind 查询之后
func (data *CarShare) AfterFind() error {
	if data.Lugg == "" {
		return nil
	}

	return json.Unmarshal([]byte(data.Lugg), &data.Luggages)
}

func CreateCarShare(sharer *CarShare) error {
	return db.Create(sharer).Error
}

func GetCarShareByDestination(destination string) ([]CarShare, int, error) {
	sharers := make([]CarShare, 10)
	var c int64
	err := db.Where("destination=?", destination).Find(&sharers).Count(&c).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return sharers, int(c), err
}

func GetCarShareById(CarShareId int) (CarShare, error) {
	var sharer CarShare
	err := db.Where("id=?", CarShareId).First(&sharer).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return sharer, err
}

func GetCarShareByUser(Userid int) ([]CarShare, int, error) {
	sharers := make([]CarShare, 10)
	var c int64
	err := db.Where("userid=?", Userid).Find(&sharers).Count(&c).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return sharers, int(c), err
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
