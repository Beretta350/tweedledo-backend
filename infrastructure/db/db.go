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
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		log.Printf("P=db M=ConnectDB env=%v connecting in postgres", env)
		dsn = os.Getenv("dsn")
		db, err = gorm.Open(postgres.Open(dsn))
	} else {
		dsn = os.Getenv("dsnTest")
		db, err = gorm.Open(sqlite.Open(dsn))
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
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
