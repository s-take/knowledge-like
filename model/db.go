package model

import (
	"github.com/jinzhu/gorm"
	// import sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "db/sample.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.SingularTable(true)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Knowledge{})
	db.AutoMigrate(&Like{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT").AddForeignKey("knowledge_id", "knowledge(id)", "RESTRICT", "RESTRICT")
}
