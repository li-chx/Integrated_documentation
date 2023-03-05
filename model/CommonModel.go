package model

import (
	"github.com/Haroxa/Integrated_documentation/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB = mysql.MysqlDb

func init() {
	// 自动迁移模式  自动创建表，添加缺少列和索引
	db.AutoMigrate(
		&User{},
	)
}
