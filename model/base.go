package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test.com/config"
)

type Base struct {
	*gorm.DB `gorm:"-" json:"-"`
	Id       int64  `gorm:"primaryKey" json:"id"`
	Ctime    string `json:"ctime"` //日期时间字段统一设置为字符串即可
	Uptime   string `json:"uptime"`
	//DeletedAt gorm.DeletedAt `json:"deltime"`   // 如果开发者需要使用软删除功能，打开本行注释掉的代码即可，同时需要在数据库的所有表增加字段deleted_at 类型为 datetime
}

func UseDbConn() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlConfig.UserName,
		config.MysqlConfig.Password,
		config.MysqlConfig.Host,
		config.MysqlConfig.Port,
		config.MysqlConfig.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db conn fail : ", err)
		panic(err)
	}
	return db
}
