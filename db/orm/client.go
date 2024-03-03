package orm

import (
	"fmt"
	"log"

	"github.com/hadanhtuan/go-sdk"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbEnv sdk.ORMEnv
)

func ConnectDB() *gorm.DB {
	sdk.ParseENV(&dbEnv)

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbEnv.Host,
		dbEnv.Port,
		dbEnv.DBName,
		dbEnv.DBUser,
		dbEnv.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("🗃️  Connected Successfully to the database")
	return db
}
