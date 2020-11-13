# Picsrc
A simple Web api implement by gin to upload and show Images.

## Api列表
### /api/pircture
- post方法
用于上传一张图片  
前端输入：
    - file:选一个图片上传
    - tag:给图片加的一个标签(现在暂时没什么用,可以跳过)
    - token:上传输入的密钥，密钥不正确不能上传文件  
后端返回：    
```json
{
    "msg": "blabla"
}
```

- get方法
获取上传的信息  
    - page:获取第几页的图片
    - page_size:每页有几张图片  
后端返回：
```json
{
    "iamges": [
        {
            "ID": 1,
            "CreatedAt": "2020-10-19T07:20:13Z",
            "UpdatedAt": "2020-10-19T07:20:13Z",
            "DeletedAt": null,
            "Url": "localhost:8080/Files/1603092013.jpg",
            "IsDelete": true,
            "Tag": "test"
        }
    ]
}
```
