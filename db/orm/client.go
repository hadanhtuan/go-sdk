package orm

import (
	"fmt"
	"log"

	"github.com/hadanhtuan/go-sdk/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectORM() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.AppConfig.ORM.Host,
		config.AppConfig.ORM.Port,
		config.AppConfig.ORM.DBName,
		config.AppConfig.ORM.DBUser,
		config.AppConfig.ORM.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("🚀 Connected Successfully to the database")
	return db
}
