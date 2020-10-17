package main

import (

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)
func getTimeStamp() (timestamp int64){
	timestamp = time.Now().Unix()
	return
}

//User upload file should never be trusted
//so replace name with timestamp
func ParseFileName(filename string) (newfilename string){
	extension := filepath.Ext(filename)
	timestamp_int := int(getTimeStamp())
	newfilename = strconv.Itoa(timestamp_int)+extension
	return
}
func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		file.Filename = ParseFileName(file.Filename)
		dir,_:=os.Getwd()
		c.SaveUploadedFile(file, dir+"/Files/"+file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}
