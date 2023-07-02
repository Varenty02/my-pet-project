package main

import (
	"MyPetProject/component/appctx"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	log.Println("server start")
	dsn := os.Getenv("MY_SQL_CONN")
	secretKey := os.Getenv("SECRETKEY")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Can connect to db")
	}
	db = db.Debug()

	appCtx := appctx.NewAppContext(db, secretKey, client)
	r := gin.Default()
	v1 := r.Group("/v1")
	setupRoute(appCtx, v1)
	r.Run(":3007")
	err = client.Close()
	if err != nil {
		log.Fatalf("Failed to close Redis connection: %v", err)
	}
}
