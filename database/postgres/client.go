package orm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectORM(host, port, dbName, username, password string) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		port,
		dbName,
		username,
		password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to Postgres")
	}

	fmt.Println("[ 🚀 ] Connected Successfully to Postgres")
	return db
}
