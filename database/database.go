package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func (db *Database) InitDatabase() {
	db.AutoMigrate(&Task{})
}

func NewDatabase() *Database {
	conn, err := gorm.Open(sqlite.Open("sidecar.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db := new(Database)
	db.DB = conn
	db.InitDatabase()
	return db
}

func (d *Database) Close() {
	d.Close()
}
