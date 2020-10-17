package util

import (
	"github.com/jinzhu/gorm"
)
type Image struct {
	gorm.Model
	Url string
	IsDelete bool
	Tag string
}
func init(){

}
