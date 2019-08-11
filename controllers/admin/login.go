package admin

import (
	"github.com/astaxie/beego"
	"simple_blog/models"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Login()  {
	if c.isPost() {
		userName := c.GetString("userName")
		password := c.GetString("password")
		admin := new(models.Admin)
		admin.UserName = userName
		if err := admin.Read("userName");err !=nil{
			c.ajaxMsg("用户名错误",201)
		}else if Md5([]byte(password)) != admin.Pwd{
			c.ajaxMsg("密码错误",202)
		}else {
			c.SetSession("admin", *admin)
			c.ajaxMsg("登录成功",200)
		}
	}
	c.TplName = "admin/login.html"
}

//登出
func (c *LoginController) LoginOut() {
	c.DelSession("admin")
	c.redirect(beego.URLFor("LoginController.Login"))
}
