package main

import (
	"github.com/astaxie/beego"
	"tianwen_blog/controllers"
	_ "tianwen_blog/models"
	_ "tianwen_blog/routers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

