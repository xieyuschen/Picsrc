# Picsrc
A simple Web api implement by gin to upload and show Images.

# Api
## 所需环境
- 使用mysql数据库进行图片信息的存储，所需信息需要在`Config.json`进行设置。
```json
//设置mysql的连接选项
//注意本地需要有名为Dbname的内容
{
    "DbSettings":{
        "Username":"root",
        "Password" :"root",
        "Hostname" :"127.0.0.1:3306",
        "Dbname" :"imagedb"
    }
}
``` 

## 如何启动后端程序
首先需要安装go环境，然后在`main.go`目录下运行`go run main.go`。当然好像这个东西不需要前端搞，需要我把它扔到服务器上面去。
## Api列表
### 1.upload
post方法，完成图像资源的上传。
所需参数：
```json
{
    //暂时未对可上传文件后缀名进行限制
    "file":"upload a file",//choose a picture to load
    "tag":"Add tag for this picture"
}
```
上传文件大小限制为8MiB
### 2. getimages
get方法，获取一定数目的图片信息
```json
{
    "page":"get the certain page based 0",
    "pagesize":"how many images in a page"
}
```
返回`pagesize`数目的图片信息，这里不返回图片本身而是一个相对url。配置一下nginx就可以用了。  