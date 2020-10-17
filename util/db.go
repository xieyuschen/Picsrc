package util

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
func Check(err error){
	if err!=nil{
		log.Panic(err)
	}
}
func init(){
	settings := DbSettings{username: "root",password: "root",hostname: "127.0.0.1:3306",dbname: "imagedb"}
	//"root:@tcp(127.0.0.1:3306)/?parseTime=true&charset=utf8"
	connStr := dsn(settings)
	msdb, err:= sql.Open("mysql",connStr)
	Check(err)
	msdb.Exec("create database if not exists "+settings.dbname+" character set utf8")
	msdb.Close()
	db, err = gorm.Open("mysql",dsn(settings))
	Check(err)
	if !db.HasTable(&ImageType){
		db.CreateTable(&ImageType)
		log.Println("Create Image table successfully")
	}
}
func AddImage(image Image){
	db.Create(&image)
}