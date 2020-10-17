package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"Picrsc/util"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", UploadFile)

	router.Run(":8080")
}
func UploadFile(c *gin.Context){
		// single file
		file, _ := c.FormFile("file")
		tag := c.PostForm("tag")
		file.Filename =util.ParseFileName(file.Filename)
		dir,_:=os.Getwd()

		c.SaveUploadedFile(file, dir+"/Files/"+file.Filename)
		image := util.Image{Url:dir+"/Files/"+file.Filename,IsDelete:true,Tag:tag}
		util.AddImage(image)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}