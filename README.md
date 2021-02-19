# Picsrc
A simple Web api implement by gin to upload and show Images.

# Api列表
## 图片
### 1. 上传图片
post方法，路径为：`/api/picture`  
- 前端参数放在body中进行传入：
    - file:选一个图片上传
    - tag:给图片加的一个标签(现在暂时没什么用,可以跳过)
    - token:上传输入的密钥，密钥不正确不能上传文件  
- 后端依据状态码判断成功与否: 
    - `400`: 文件过大
    - `403`: 上传文件的token出错
    - `500`: 服务器上传文件失败 
    - `200`：上传文件成功，将返回如下信息：  
```json
{
    "filename":"文件名称",
    "id":"根据该id可以获取到图片"
}
```
### 2. 获取图片列表
get方法，路径为：`/api/picture`  
- 前端传入参数放在url中，  
    - page:获取第几页的图片
    - page_size:每页有几张图片
    - tag:标记该图片属于哪个内容，如果为`工作使用`则传入空，若为娱乐使用则传入`playground`
get请求的url形如：`127.0.0.1:8080/api/picture?page=0&page_size=5&tag=test`  
- 后端返回：
```json
{
    "iamges": [
        {
            "ID": 1,
            "CreatedAt": "2020-10-19T07:20:13Z",
            "UpdatedAt": "2020-10-19T07:20:13Z",
            "DeletedAt": null,
            "Url": "localhost:8080/Files/1603092013.jpg",
            "Tag": "test"
        }
    ]
}
```
### 删除图片
delete方法，请求路径形如`/api/picture/<image_id>`
例如删除id为3的图片，则发送delete请求，路径为`/api/picture/3`  
- 前端传入
    - url中放入要删除的id
- 后端返回
后端返回状态码标识对应的情况。  
    - `400`:前端传入的id非法
    - `404`:没有找到传入图片的id
    - `200`:成功
    