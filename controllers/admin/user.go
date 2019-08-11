package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"simple_blog/models"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

func (c *UserController) List() {
	page,pagesize := c.getPage()
	user := new(models.User)
	var users []*models.User
	user.Query().Limit(pagesize,(page-1)*pagesize).All(&users)
	if c.Ctx.Request.Method =="POST"{
		c.Data["json"] = users
		c.ServeJSON()
	}
	c.Data["list"] = users
	c.Data["total"],_ = user.Query().Count()
	c.display("admin/user/list")
}

func (c *UserController) Edit()  {
	id,_ := c.GetInt("id")
	user := new(models.User)
	user.Id = id
	user.Read()
	if c.Ctx.Request.Method == "POST"{
		Avatar := c.GetString("Avatar")
		UserName := c.GetString("UserName")
		IsSuper,_ := c.GetBool("IsSuper")
		Status,_ := c.GetInt("Status")
		RealName := c.GetString("RealName")
		Mobile := c.GetString("Mobile")
		Email := c.GetString("Email")
		Sex := c.GetString("Sex")
		Birthday := c.GetString("Birthday")
		user.Avatar = Avatar
		user.UserName = UserName
		user.IsSuper = IsSuper
		user.Status = Status
		user.RealName = RealName
		user.Mobile = Mobile
		user.Email = Email
		user.Sex = Sex
		BirthdayTime, _ := time.ParseInLocation("2006-01-02", Birthday, time.Local)
		user.Birthday = BirthdayTime
		re := make(map[string]interface{})
		re["code"] = 200
		re["msg"] = "修改成功"
		if _,err := user.Update();err != nil{
			re["code"] = 200
			re["msg"] = "修改失败"
			re["err"] = err.Error()
		}
		c.Data["json"] = re
		c.ServeJSON()
	}
	user.UserPwd = ""
	c.Data["user"] = user
	c.display("admin/user/edit")
}

func (c *UserController) EditPwd()  {
	re := make(map[string]interface{})
	re["code"] = 200
	re["msg"] = "添加成功"
	id,_ := c.GetInt("id")
	UserPwd := strings.Replace(c.GetString("UserPwd"), " ", "", -1)
	NewPwd := strings.Replace(c.GetString("NewPwd"), " ", "", -1)
	VisiblePwd := strings.Replace(c.GetString("VisiblePwd"), " ", "", -1)
	if id == 0{
		re["code"] = 205
		re["msg"] = "ID不能为空"
		c.Data["json"] = re
		c.ServeJSON()
	}
	if UserPwd == ""{
		re["code"] = 201
		re["msg"] = "原密码不能为空"
		c.Data["json"] = re
		c.ServeJSON()
	}
	if NewPwd == ""{
		re["code"] = 202
		re["msg"] = "新密码不能为空"
		c.Data["json"] = re
		c.ServeJSON()
	}
	if NewPwd != VisiblePwd {
		re["code"] = 204
		re["msg"] = "两次密码不一致"
		c.Data["json"] = re
		c.ServeJSON()
	}
	user := new(models.User)
	user.Id = id
	user.Read()
	if user.UserPwd != Md5([]byte(UserPwd)) {
		re["code"] = 207
		re["msg"] = "原密码错误"
		c.Data["json"] = re
		c.ServeJSON()
	}
	user.UserPwd = Md5([]byte(NewPwd))
	if _, err := user.Update("UserPwd"); err != nil{
		re["code"] = 206
		re["msg"] = "修改密码失败"
		re["err"] = err.Error()
		c.Data["json"] = re
		c.ServeJSON()
	}
	c.Data["json"] = re
	c.ServeJSON()
}

func (c *UserController) Add() {
	if c.Ctx.Request.Method == "POST"{
		Avatar := c.GetString("Avatar")
		UserName := strings.Replace(c.GetString("UserName"), " ", "", -1)
		UserPwd1 := strings.Replace(c.GetString("UserPwd1"), " ", "", -1)
		UserPwd2 := strings.Replace(c.GetString("UserPwd2"), " ", "", -1)
		IsSuper,_ := c.GetBool("IsSuper")
		Status,_ := c.GetInt("Status")
		RealName := c.GetString("RealName")
		Mobile := c.GetString("Mobile")
		Email := c.GetString("Email")
		Sex := c.GetString("Sex")
		Birthday := c.GetString("Birthday")
		re := make(map[string]interface{})
		re["code"] = 200
		re["msg"] = "添加成功"
		if UserName == "" {
			re["code"] = 201
			re["msg"] = "会员名称为空"
			c.Data["json"] = re
			c.ServeJSON()
		}
		if UserPwd1 == "" {
			re["code"] = 201
			re["msg"] = "密码为空"
			c.Data["json"] = re
			c.ServeJSON()
		}
		if UserPwd1 != UserPwd2{
			re["code"] = 204
			re["msg"] = "两次密码不一致"
			c.Data["json"] = re
			c.ServeJSON()
		}
		user := new(models.User)
		user.Avatar = Avatar
		user.UserName = UserName
		user.UserPwd = Md5([]byte(UserPwd1))
		user.IsSuper = IsSuper
		user.Status = Status
		user.RealName = RealName
		user.Mobile = Mobile
		user.Email = Email
		user.Sex = Sex
		BirthdayTime, _ := time.ParseInLocation("2006-01-02", Birthday, time.Local)
		user.Birthday = BirthdayTime
		if _,err := user.Insert();err != nil{
			re["code"] = 205
			re["msg"] = err.Error()
			c.Data["json"] = re
			c.ServeJSON()
		}
		c.Data["json"] = re
		c.ServeJSON()
	}
	c.display("admin/user/add")
}

func (c *UserController) AvatarUpload()  {
	uploadUrl := beego.AppConfig.String("uploadUrl")
	_, h, _ := c.GetFile("avatar")
	file := "static/upload/" + h.Filename
	var obj interface{}
	req:=httplib.Post(uploadUrl)
	req.PostFile("file",file)//注意是全路径
	req.Param("output","json")
	req.Param("scene","")
	req.Param("path","")
	req.ToJSON(&obj)
	c.Data["json"] = obj
	c.ServeJSON()
}


