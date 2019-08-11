package admin

import (
	"simple_blog/models"
	"strings"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) Profile() {
	adminInfo := c.GetSession("admin")
	admin, _ := adminInfo.(models.Admin)
	c.Data["admin"] = admin
	c.display("admin/admin/profile")
}

func (c *AdminController) Settings() {
	adminInfo := c.GetSession("admin")
	admin, _ := adminInfo.(models.Admin)
	if c.isPost(){
		Avatar := c.GetString("Avatar")
		UserName := c.GetString("UserName")
		RealName := c.GetString("RealName")
		Mobile := c.GetString("Mobile")
		Email := c.GetString("Email")
		AdminModel := new(models.Admin)
		AdminModel.Id = admin.Id
		AdminModel.Avatar = Avatar
		AdminModel.UserName = UserName
		AdminModel.RealName = RealName
		AdminModel.Tel = Mobile
		AdminModel.Email = Email
		if err := AdminModel.Update("Avatar","UserName","RealName","Tel","Email");err != nil{
			c.ajaxMsg(err.Error(),201)
		}
		c.ajaxMsg("修改成功",200)

	}
	adminData := new(models.Admin)
	adminData.Id = admin.Id
	adminData.Read()
	c.Data["admin"] = adminData
	c.display("admin/admin/settings")
}

func (c *AdminController) EditPwd() {
	adminInfo := c.GetSession("admin")
	admin, _ := adminInfo.(models.Admin)
	if c.isPost(){
		adminModel := new(models.Admin)
		UserPwd := strings.Replace(c.GetString("UserPwd"), " ", "", -1)
		NewPwd := strings.Replace(c.GetString("NewPwd"), " ", "", -1)
		VisiblePwd := strings.Replace(c.GetString("VisiblePwd"), " ", "", -1)
		if UserPwd == ""{
			c.ajaxMsg("原密码不能为空",201)
		}
		if NewPwd == ""{
			c.ajaxMsg("新密码不能为空",202)
		}
		if NewPwd != VisiblePwd {
			c.ajaxMsg("两次密码不一致",204)
		}
		adminModel.Id = admin.Id
		adminModel.Read()
		if adminModel.Pwd != Md5([]byte(UserPwd)){
			c.ajaxMsg("原密码错误",205)
		}
		adminModel.Pwd = Md5([]byte(NewPwd))
		if err := adminModel.Update("Pwd"); err != nil{
			c.ajaxMsg("修改失败",206)
		}
		c.ajaxMsg("修改成功",200)
	}
	c.display("admin/admin/editPwd")
}
