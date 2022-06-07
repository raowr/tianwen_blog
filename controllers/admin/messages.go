package admin

import (
	"tianwen_blog/models"
)

type MessagesController struct {
	BaseController
}

func (c *MessagesController) List() {
	page,pagesize := c.getPage()
	var Messages []*models.Messages
		Message := new(models.Messages)
	var lists []map[string]interface{}
	Message.Query().Limit(pagesize,(page -1 )*pagesize).All(&Messages)
	for _, value := range Messages{
		list := make(map[string]interface{})
		if value.User != nil{
			value.User.Read()
			list["UserName"] = value.User.UserName
		}else{
			list["UserName"] = value.Name
		}
		list["Content"] = value.Content
		list["Reply"] = value.Reply
		list["CreateTime"] = value.CreateTime
		list["Open"] = value.Open
		lists = append(lists, list)
	}
	if c.Ctx.Request.Method == "POST"{
		c.Data["json"] = lists
		c.ServeJSON()
	}
	c.Data["list"] = lists
	c.Data["total"],_ = Message.Query().Count()
	c.display("admin/messages/list")
}

