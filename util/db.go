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
	Username string `json:username`
	Password string	`json:password`
	Hostname string `json:hostname`
	Dbname   string `json:dbname`
}
var ImageType Image
var db *gorm.DB
func dsn(settings DbSettings) string {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	// Add ?parseTime=true
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8", settings.Username,settings.Password, settings.Hostname,settings.Dbname)
}

func init(){
	settings := DbSettings{Username: "root", Password: "root", Hostname: "127.0.0.1:3306", Dbname: "imagedb"}
	//"root:@tcp(127.0.0.1:3306)/?parseTime=true&charset=utf8"
	connStr := dsn(settings)
	msdb, err:= sql.Open("mysql",connStr)
	Check(err)
	msdb.Exec("create database if not exists "+settings.Dbname +" character set utf8")
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
func GetImages(page int,pagesize int)(images []Image){
	db.Limit(pagesize).Offset(page*pagesize).Find(&images)
	return
}