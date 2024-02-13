package orm

import (
	"fmt"
	"github.com/hadanhtuan/go-sdk/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dbConfig config.DBOrm) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.DBUser,
		dbConfig.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
