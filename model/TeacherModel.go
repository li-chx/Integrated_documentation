package model

import (
	//"fmt"
	"gorm.io/gorm"
)

type Teacher struct {
	Id        int    `json:"id" gorm:"id;primary_key;auto_increment"`
	Name      string `json:"name"`
	Course    string `json:"course"`
	Callname  string `json:"callname"`
	Cn        string `json:"cn,omitempty"`
	Homework  string `json:"homework"`
	Hw        string `json:"hw,omitempty"`
	Mark      string `json:"mark"`
	Mk        string `json:"mk,omitempty"`
	Finishway string `json:"finishway"`
	Fw        string `json:"fw,omitempty"`
	Score     string `json:"score"`
	Sc        string `json:"sc,omitempty"`
}

var Callname = map[string]interface{}{
	"总是": 0, "经常": 0, "偶尔": 0, "从不": 0, "sum": 0,
}
var CNorder = map[int]string{
	1: "总是", 2: "经常", 3: "偶尔", 4: "从不",
}

var Homework = map[string]interface{}{
	"很多": 0, "较多": 0, "正常": 0, "较少": 0, "没有": 0, "sum": 0,
}
var HWorder = map[int]string{
	1: "很多", 2: "较多", 3: "正常", 4: "较少", 5: "没有",
}

var Mark = map[string]interface{}{
	"多为90+": 0, "多为80+": 0, "多为70+": 0, "多为60+": 0, "sum": 0,
}
var MKorder = map[int]string{
	1: "多为90+", 2: "多为80+", 3: "多为70+", 4: "多为60+",
}

var Finishway = map[string]interface{}{
	"闭卷考试": 0, "开卷考试": 0, "大作业": 0, "结课论文": 0, "sum": 0,
}
var FWorder = map[int]string{
	1: "闭卷考试", 2: "开卷考试", 3: "大作业", 4: "结课论文",
}

var Score = map[string]interface{}{
	"score": 0, "sum": 0,
}

func CreateTeacher(ta *Teacher) error {
	return db.Create(ta).Error
}

func GetTeacherByNAndC(name string, course string) ([]Teacher, int, error) {
	var ta []Teacher
	temp := db
	var count, c1, c2, flag int64
	if name != "" {
		flag++ //fmt.Println("name : ", ta)
		temp = temp.Where("name=?", name).Find(&ta).Count(&c1)
		count += c1 //fmt.Println(ta)
	}
	if course != "" {
		flag++ //fmt.Println("course : ", ta)
		temp = temp.Where("course=?", course).Find(&ta).Count(&c2)
		count += c2 //fmt.Println(ta)
	}
	if flag == 2 {
		count = c2
	}

	return ta, int(count), temp.Error
}

func GetTeacherById(TeacherId int) (Teacher, error) {
	var ta Teacher
	err := db.Where("id=?", TeacherId).First(&ta).Error //找到用户并赋值
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return ta, err
}

func GetAllTeacher() ([]Teacher, int, error) {
	tas := make([]Teacher, 1000) //ta模型切片
	var c int64
	err := db.Find(&tas).Count(&c).Error //找到所有用户并赋值
	//Count函数，直接返回查询匹配的行数。
	return tas, int(c), err
}

func UpdateTeacher(ta *Teacher, mp interface{}) error {
	return db.Model(ta).Updates(mp).Error
}

func DeleteTeacher(ta Teacher) error {
	err := db.Unscoped().Delete(&ta).Error
	return err
}
