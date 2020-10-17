package util

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)
type Image struct {
	gorm.Model
	Url string
	IsDelete bool
	Tag string
}
type DbSettings struct{
	username string
	password string
	hostname string
	dbname string
}
var ImageType Image
var db *gorm.DB
func dsn(settings DbSettings) string {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	// Add ?parseTime=true
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8", settings.username,settings.password, settings.hostname,settings.dbname)
}
func init(){
	settings := DbSettings{username: "root",password: "root",hostname: "127.0.0.1:3306",dbname: "imagedb"}
	//"root:@tcp(127.0.0.1:3306)/?parseTime=true&charset=utf8"
	connStr := dsn(settings)
	msdb, _:= sql.Open("mysql",connStr)
	msdb.Exec("create database if not exists "+settings.dbname+" character set utf8")
	msdb.Close()
	db, _= gorm.Open("mysql",dsn(settings))
	if !db.HasTable(&ImageType){
		db.CreateTable(&ImageType)
		log.Println("Create Image table successfully")
	}
}
func AddImage(image Image){
	db.Create(&image)
}