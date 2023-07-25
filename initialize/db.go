package initialize

import (
	"chat/constant"
	"chat/global"
	"chat/models"
	"chat/utils"
	"fmt"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	utils.Logger(constant.Info, map[string]string{"mes": "table init"})

	// add your models to migrate into database
	tables := []interface{}{
		&models.Group{},
		&models.User{},
	}
	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			db.Migrator().CreateTable(table)
		}
	}

	utils.Logger(constant.Info, map[string]string{"mes": "table init ok"})

}

func initData(db *gorm.DB) {

	utils.Logger(constant.Info, map[string]string{"mes": "data init"})

	now := time.Now()

	db.Where(models.Group{Name: "will"}).FirstOrCreate(&models.Group{
		Name: "will",
		Img:  "www.baidu.com",
	})
	
	db.Where(models.User{Name: "Jinzhu"}).FirstOrCreate(&models.User{
		Name:           "will",
		Sex:            true,
		Bir:            now,
		Img:            "www.baidu.com",
		Phone:          "11111111111",
		UserID:         "will",
		RegisteredTime: now,
		LoginTime:      now,
		LogoutTime:     now,
	})

	utils.Logger(constant.Info, map[string]string{"mes": "data init ok"})
}

func InitDB() {
	utils.Logger(constant.Info, map[string]string{"mes": "db init"})

	mysqlInfo := global.Settings.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlInfo.User, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.DataBase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mysqlInfo.TablePrefix,
			SingularTable: true,
		},
	})
	utils.Error(err, "failed to connect database")
	global.Db = db

	// migrate sql
	migrate(db)
	// init data
	initData(db)
	utils.Logger(constant.Info, map[string]string{"mes": "db init ok"})

}
