package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zeimedee/sre_bootcamp/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type database struct {
	HOST     string
	USER     string
	PASSWORD string
	NAME     string
}

func NewDb() *database {
	db := &database{
		HOST:     os.Getenv("POSTGRES_HOST"),
		USER:     os.Getenv("POSTGRES_USER"),
		PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		NAME:     os.Getenv("POSTGRES_DB"),
	}
	return db
}

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func ConnectDb() {

	er := godotenv.Load()
	if er != nil {
		log.Fatalf("Error loading .env file: %v", er)
	}

	Db := NewDb()
	dsn := fmt.Sprintf("host='%s' user='%s' password='%s' dbname='%s' port=5432 sslmode=disable TimeZone=Asia/Shanghai", Db.HOST, Db.USER, Db.PASSWORD, Db.NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Student{})

	DB = DBInstance{
		Db: db,
	}

}
