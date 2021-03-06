package main

import (
	_ "github.com/Tinywan/webcodec/routers" // _ 表示只调用routers的 init()方法， 这个包只引入执行了里面的 init 函数
	"github.com/astaxie/beego"
	"github.com/Tinywan/webcodec/models"
	"net/http"
	"html/template"
)

const VERSION = "1.0.0"

func main() {
	models.Init();

	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	beego.AppConfig.Set("version", VERSION)

	// 路由设置
	//beego.Router("/", &controllers.MainController{}, "*:Index")
	//beego.Router("/login", &controllers.MainController{}, "*:Login")
	//beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	//beego.Router("/profile", &controllers.MainController{}, "*:Profile")
	//beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
	// beego.Run 执行之后，我们看到的效果好像只是监听服务端口这个过程，但是它内部做了很多事情
}

