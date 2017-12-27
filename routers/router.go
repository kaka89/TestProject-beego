package routers

import (
	"TestProject-beego/application/rest"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
)

/**
 *	在beego里面，路由的配置方式有："固定路由"、"正则路由"、"自定义"、"自动匹配"、"注解路由"、"namespace"等几种方式，互相之间不冲突，可以交叉使用.
 *  不过这个测试项目里面只使用namespace的方式，原因有：NS除了能够支持其他几种的配置之外，还能够额外的定义一些处理，例如前置，后置等，故后面均采用NS的方式进行路由的配置.
 *  所有的路由均采用小写的方式进行命名
 */
func init() {

	ns := beego.NewNamespace("/v1", //所有API都以"/v1"起始，做API版本控制之用
		beego.NSCond(func(ctx *context.Context) bool { //对请求做一些简单公共的判断，判断正确才进入之后的请求
			if ctx.Input.Domain() == "localhost" || ctx.Input.Domain() == "api.umasuo.com" {
				return true
			}
			return false
		}),
		//beego.NSBefore(auth),该NS下所有API都需要进入
		//beego.NSNamespace("/object", // 子namespace
		//	beego.NSInclude( //这种方法是直接包含了所有通过注解暴露的方法
		//		&rest.ObjectController{},
		//	),
		//	beego.NSRouter("/", &rest.ObjectController{}, "get:GetAll"),       //基本配置方法
		//	beego.NSRouter("/:objectId", &rest.ObjectController{}, "get:Get"), //基本配置方法
		//),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&rest.UserController{},
			),
			beego.NSRouter("/", &rest.UserController{}, "post:Post"),
		),
		//beego.NSAfter(auth),
	)

	//添加进入beego
	beego.AddNamespace(ns)
}
