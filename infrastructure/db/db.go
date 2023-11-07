package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/tweedledo/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	log.Printf("P=db M=init basepath=%v", basepath)

	err := godotenv.Load(basepath + "../../.env")
	if err != nil {
		log.Printf("Error loading .env files")
	}
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		log.Printf("P=db M=ConnectDB env=%v trying to connect in postgres", env)
		if len(os.Getenv("POSTGRES_URL")) > 0 {
			dsn = os.Getenv("POSTGRES_URL")
		} else {
			dsn = os.Getenv("dsn")
		}
		db, err = gorm.Open(postgres.Open(dsn))
	} else {
		log.Printf("P=db M=ConnectDB env=%v trying to connect in memory sqlite", env)
		dsn = os.Getenv("dsnTest")
		db, err = gorm.Open(sqlite.Open(dsn))
	}

	if err != nil {
		panic("Error connecting to database:" + err.Error())
	}

	if os.Getenv("debug") == "true" {
		db.Config.Logger = logger.Default.LogMode(logger.Warn)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		log.Printf("P=db M=ConnectDB env=%v auto migrating", env)
		db.AutoMigrate(&domain.TaskList{}, &domain.Task{})
	}

	log.Printf("P=db M=ConnectDB env=%v database connected", env)
	return db
}
