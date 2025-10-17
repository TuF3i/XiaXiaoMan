package dao

import (
	"XiaXiaoMan/core"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDataBase() {
	//构造MySQL URL
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		core.Conf.MySQLUser,
		core.Conf.MySQLPasswd,
		core.Conf.MySQLIPAddr,
		core.Conf.MySQLPort,
		core.Conf.MySQLDB)

	//连接数据库
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connect to MySQL Fail!!!\n")
		os.Exit(1)
	}

	////迁移数据库
	//err = db.AutoMigrate(&models.QAData{}, &models.User{})
	//if err != nil {
	//	return fmt.Errorf("数据库迁移失败:%v", err)
	//}

	core.DB = db
}
