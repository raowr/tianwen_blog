package main

import (
	_ "simple_blog/routers"
	"github.com/astaxie/beego"
	_"simple_blog/models"
	"simple_blog/controllers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

