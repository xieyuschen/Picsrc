package main

import (
	"os"
	"github.com/gin-contrib/cors"
	"Picrsc/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 17  // 8 MiB
	pic:= router.Group("/api/picture")
	{
		pic.Use(cors.Default())
		pic.POST("", UploadFile)
		pic.GET("",GetImages)
		pic.GET("/hello",Helloworld)
		pic.DELETE("/:image_id",DeleteImages)
	}
	router.Run(":8080")
}
func UploadFile(c *gin.Context){
	file, err := c.FormFile("file")
	if err!=nil{
		fmt.Println(err.Error())
	}
	tag := c.PostForm("tag")
	token:=c.PostForm("token")
	if token!=util.Settings.Token{
		c.JSON(http.StatusForbidden,gin.H{
			"msg":"Token不正确，无法上传图片",
		})
		return
	}
	if (file.Size>5*(2 << 20)){
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":"文件太大",
			})
			return
		}
	str, _ := os.Getwd()
	fmt.Println(str)

	file.Filename =util.ParseFileName(file.Filename)
	err=c.SaveUploadedFile(file, "../Picsrc/Files/"+file.Filename)

	image := util.Image{Url:"146.56.199.136/Files"+file.Filename,IsDelete:false,Tag: tag}
	util.AddImage(&image)
	c.JSON(http.StatusOK, gin.H{
		"filename":file.Filename,
		"id":image.ID,
	})
	return
}
func GetImages(c *gin.Context){
	page,_ :=strconv.Atoi( c.Query("page"))
	pagesize,_ := strconv.Atoi(c.Query("page_size"))
	tag := c.Query("tag")
	images := util.GetImages(page,pagesize,tag)
	c.JSON(http.StatusOK,gin.H{
		"images":images,
	})
}
func Helloworld(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg":"helloworld",
	})
}
func DeleteImages(ctx *gin.Context){
	image_id,err:=strconv.Atoi(ctx.Param("image_id"))
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{})
		return
	}
	err=util.DeleteImage(image_id)
	if err!=nil{
		ctx.JSON(http.StatusNotFound,gin.H{})
		return
	}else {
		ctx.JSON(http.StatusOK,gin.H{})
		return
	}
}
