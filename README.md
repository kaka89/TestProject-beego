# BeeGo 框架测试项目， API 项目，非Web项目
测试工程 //TODO 待完善


# 项目基本规约
项目在开发过程中需要遵守一些基本的开发规约，既方便自己当前的开发，也方便后来人理解项目.
## 项目地址
默认情况下，go项目均放在默认目录:`/Users/your-computer-user-name/go/src/`，例如我的是`/Users/umasuo/go/src/`.

`PS. 上面目录是在Mac下的目录，如果开发机是windows或者Linux等，请自行设置，并补充进来`

## 项目目录结构
项目结构是根据个人以前经验和习惯组织，可能
```
- ProjectName           ---- 整个工程的根目录
    -application        ---- 应用层根目录，应用层的作用是对外暴露服务
        -dto            ---- 对外暴露服务的数据结构
            -builder    ---- 数据结构的builder
        -rest           ---- REST接口，目前系统对外通过REST API的方式暴露服务，后期如果需要，可通过RPC等方式暴露，到时候可再增加
        -service        ---- 应用层服务包，于独立领域无关的服务
    -conf               ---- 系统配置
    -domain             ---- 独立领域层，具体业务服务
        -entity         ---- 领域实体，基本于数据库表结构一一对应
        -service        ---- 领域服务，独立不交叉
    -infrastructure     ---- 基础组件包，于具体业务无关、工具类、通用类的组件，均防止在这里面
        -auth           ---- 权限认证
        -repository     ---- 封装数据库的操作,提供给domain-service使用
        -utils          ---- 工具类
        -···            ---- 其他
    -routers            ---- 路由配置包，目前只能够放置在这里
    -tests              ---- 单元测试目录
    main.go             ---- 项目的Main 包，用于启动项目
```
## 包引用
- 外部包引用按照标准来，例如："github.com/astaxie/beego"
- 内部包引用，采用以项目目录为根目录的方式来引用，例如："LMS/routers"，而不是相对路径
## 基本开发流程
    定义API（目前通过API形式暴露服务）--> 定义数据库 --> 开发具体业务

`ps.如果API有变动，需预先通知其他相关团队和服务`

## 提交基本规约
1. 提交之前不能够block项目的build，本地先build完再提交
2. 重大更改不能够提交到master，以分支的形式开发，等待测试完成之后，再merge进master
3. 版本上线需开分支，如需fix线上bug，则提交到对应分之，后面再merge进master

## 代码基本规约
一些典型的代码的规约
### router
### REST API Controller
### dto
### entity
### application service
### domain service

 
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


