package db

import (
	"log"
	"os"
	"time"
	"yugod-backend/app/config"
	"yugod-backend/app/dao"
	"yugod-backend/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	gormConfig := &gorm.Config{
		// 表名配置：添加前缀
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ysy_",
			SingularTable: true,
		},
	}
	// mysql Debug
	if config.App.Debug {
		gormLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // Disable color
			},
		)
		gormConfig.Logger = gormLogger
	}

	db, err := gorm.Open(mysql.Open(config.DB.GetDSN()), gormConfig)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	// 传入dao层
	dao.DB = db
	err = db.AutoMigrate(
		&model.ClickVolume{},
		&model.User{},
	)

	if err != nil {
		log.Fatalf("Got error when migrate database, the error is '%v'", err)
	}
}
