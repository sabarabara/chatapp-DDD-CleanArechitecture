package config

import (
	"fmt"
	"os"
)

func InitDB()(string){

	//.envからの読み込み
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSL_MODE")
	tiemzone := os.Getenv("DB_TIME_ZONE")

	dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        host, user, password, dbname, port, sslmode, tiemzone,
    )


	return dsn
}