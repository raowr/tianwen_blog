package admin

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type BaseController struct {
	beego.Controller
}

// Prepare implemented Prepare method for baseRouter.
func (c *BaseController) Prepare() {
	// flash := beego.NewFlash()
	// Setting properties.
	controllerName, _ := c.GetControllerAndAction()
	c.StartSession()
	c.Data["PageStartTime"] = time.Now()
	user := c.GetSession("admin")
	if user == nil && controllerName != "LoginController" {
		c.redirect(beego.URLFor("LoginController.Login"))
	}

}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

func (c *BaseController) display(tpl string)  {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["head"] = "admin/head.html"
	if len(tpl) == 0{
		c.TplName = "admin/info.html"
	} else {
		c.TplName = tpl+".html"
	}
	c.Layout = "admin/layout.html"
}

func (c *BaseController) getTime() time.Time {
	timezone := float64(0)
	add := timezone * float64(time.Hour)
	return time.Now().UTC().Add(time.Duration(add))
}

func (c *BaseController) getPage() (page,pagesize int) {
	var err error
	if page, err = strconv.Atoi(c.GetString("page")); err != nil || page < 1 {
		page = 1
	}
	if pagesize, err = strconv.Atoi(c.GetString("pagesize")); err != nil || pagesize < 1{
		pagesize = 5
	}
	return page,pagesize
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
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

//ajax返回 列表
func (c *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	c.Data["json"] = out
	c.ServeJSON()
	c.StopRun()
}

// 是否POST提交
func (c *BaseController) isPost() bool {
	return c.Ctx.Request.Method == "POST"
}