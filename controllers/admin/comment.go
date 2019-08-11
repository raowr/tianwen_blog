package admin

import (
	"simple_blog/models"
)

type CommentController struct {
	BaseController
}

func (c *CommentController) List() {
	page,pagesize := c.getPage()
	var Comments []*models.Comment
	Comment := new(models.Comment)
	Article := new(models.Article)
	Album := new(models.Album)
	var lists []map[string]interface{}
	Comment.Query().Limit(pagesize,(page -1 )*pagesize).All(&Comments)
	for _, value := range Comments{
		list := make(map[string]interface{})
		Title := ""
		if value.Type == 1 {
			Article.Id = value.FirstId
			Article.Read()
			Title = Article.Title
		}else if value.Type == 2{
			Album.Id = value.FirstId
			Album.Read()
			Title = Album.Title
		}
		if value.User != nil{
			value.User.Read()
			list["UserName"] = value.User.UserName
		}else{
			list["UserName"] = value.Name
		}
		list["Content"] = value.Content
		list["Title"] = Title
		list["CreateTime"] = value.CreateTime
		lists = append(lists, list)
	}
	if c.Ctx.Request.Method == "POST"{
		c.Data["json"] = lists
		c.ServeJSON()
	}
	c.Data["list"] = lists
	c.Data["total"],_ = Comment.Query().Count()
	c.display("admin/comment/list")
}