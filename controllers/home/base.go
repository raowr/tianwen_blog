package home

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) display(tpl string)  {
	c.TplName = tpl+".html"
	c.Layout = "home/layout.html"
}

//ajax返回
func (c *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	c.Data["json"] = out
	c.ServeJSON()
	c.StopRun()
}

