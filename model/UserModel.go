package model

import "gorm.io/gorm"

// 要保存到数据库的模型
// ,omitempty 加在 json 处才会不显示空值

type User struct { // 个人信息
	Id       int    `json:"id" gorm:"id;primary_key;auto_increment"`
	Username string `json:"username,omitempty" gorm:"username"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password,omitempty" gorm:"password"`
	Name     string `json:"name,omitempty" gorm:"name"`
	Idcard   string `json:"idcard,omitempty" gorm:"student_number"`
	Sex      string `json:"sex,omitempty" gorm:"sex"`
	School   string `json:"school,omitempty" gorm:"school"`
	Major    string `json:"major,omitempty" gorm:"major"`
	QQ       string `json:"qq,omitempty" gorm:"qq"`
	Wechat   string `json:"wechat,omitempty" gorm:"wechat"`
	Phone    string `json:"phone,omitempty" gorm:"phone"`
}

// 不保存的模型

type UserLogin struct { // 登录信息
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) error {
	//在数据库中创建一个新用户
	return db.Create(user).Error
}

func FindUser(email string) (User, error) {
	var user User
	//在数据库中查找用户名，找到就给user赋值,找不到就为空
	//正常注册时会出现 record not found
	//First() 函数找不到record的时候，会返回error: record not found ，
	//而Find() 则是返回nil
	err := db.Where("email=?", email).First(&user).Error
	return user, err
}

func GetUserById(UserId int) (User, error) {
	var user User
	err := db.Where("id=?", UserId).First(&user).Error //找到用户并赋值
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return user, err
}

func UpdateUser(user User, mp interface{}) error {
	//对指定的用户模型，更新多个字段值，并判断是否有误
	return db.Model(&user).Updates(mp).Error
}

func GetAllUser() ([]User, int, error) {
	users := make([]User, 1000) //user模型切片
	var c int64
	err := db.Find(&users).Count(&c).Error //找到所有用户并赋值
	//Count函数，直接返回查询匹配的行数。
	return users, int(c), err
}

func DeleteUser(user User) error {
	//删除指定用户以及数据库中的删除记录
	err := db.Unscoped().Delete(&user).Error
	return err
}
