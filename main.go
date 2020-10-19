package main

import (
	"Picrsc/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", UploadFile)
	router.GET("/getimages",GetImages)
	router.GET("/",Helloworld)
	router.Run(":8080")
}
func UploadFile(c *gin.Context){
		// single file
		file, err := c.FormFile("file")
		util.Check(err)
		tag := c.PostForm("tag")
		file.Filename =util.ParseFileName(file.Filename)

		c.SaveUploadedFile(file, "/Files/"+file.Filename)
		image := util.Image{Url:"/Files/"+file.Filename,IsDelete:true,Tag:tag}
		util.AddImage(image)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
func GetImages(c *gin.Context){
	page,_ :=strconv.Atoi( c.Query("page"))
	pagesize,_ := strconv.Atoi(c.Query("pagesize"))
	images := util.GetImages(page,pagesize)
	c.JSON(http.StatusOK,gin.H{
		"iamges":images,
	})
}
func Helloworld(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg":"helloworld",
	})
}