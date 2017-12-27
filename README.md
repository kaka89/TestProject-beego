# BeeGo 框架测试项目， API 项目，非Web项目
测试工程

#项目基本规约
## 如何引用项目内部包
以项目目录为根目录，引用采用`绝对路径`方式应用，不采用`相对路径`
 
# 如何运行
## 独立运行
在项目目录下面直接运行
```
go run main.go
```
或者
```shell
bee run
```

## 打包成docker镜像运行

1、打包成docker镜像
```shell
godep save
docker build . -t testproject:1.0
```
2、运行
```shell
docker run -p 8080:8080 --name testproject -d testproject:1.0
```

`备注：记得运行'godep save',否则会出现此类错误；`[godep: No Godeps found (or in any parent directory)](https://stackoverflow.com/questions/47975830/when-build-beego-docker-image-with-default-docker-file-show-error-godep-no-g)


