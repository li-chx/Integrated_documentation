package mysql

import (
	"fmt"
	"github.com/Haroxa/Integrated_documentation/config"
	"gorm.io/gorm"
)

var MysqlDb *gorm.DB
var MysqlDbErr error

func init() {
	//  获取  数据库等 相关配置
	dbConfig := config.GetDbConfig()
	//  打印 ， 并返回字符串
	dbDSN := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%v&loc=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Hostname,
		dbConfig.Port,
		dbConfig.Dbname,
		dbConfig.Charset,
		dbConfig.ParseTime,
		dbConfig.Local,
	)
	//  打开数据库
	MySqlDb, MySqlDbErr = gorm.Open(mysql.Open(dbDSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	//  返回错误
	if MySqlDbErr != nil {
		panic("database open error" + MySqlDbErr.Error())
	}
}
