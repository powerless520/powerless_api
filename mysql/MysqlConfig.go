package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var mysqlEngine *gorm.DB

func Init() *gorm.DB {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Asphalt1018.@tcp(127.0.0.1:3306)/api_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Println("mysql connect err:" + err.Error())
		return nil
	}

	mysqlEngine = db
	// todo sql配置
	return mysqlEngine
}

func GetMysqlEngine() *gorm.DB  {
	return mysqlEngine
}