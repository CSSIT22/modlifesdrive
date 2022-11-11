package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type FileInfo struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Size      int64
	Extention string
}

var DB *gorm.DB

func SetUpDb() error {
	db, err := gorm.Open(sqlite.Open("./db/fileSystem.db"), &gorm.Config{})
	DB = db
	if err == nil {
		db.AutoMigrate(&FileInfo{})
	}
	return err
}
