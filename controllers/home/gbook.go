package home

import (
	"math"
	"tianwen_blog/models"
	"time"
)

type list struct {
	Name string
	Avatar string
	CreateTime string
	Content string
	Reply string
}
type GbookController struct {
	BaseController
}

func (c *GbookController) Index() {
	//留言列表
	message := new(models.Messages)
	var messages []*models.Messages
	var lists []*list
	if c.Ctx.Request.Method == "POST"{
		page,_ := c.GetInt("page")
		message.Query().Filter("Open",1).Limit(4,(page-1)*4).All(&messages)
		lists = c.getMessag(messages)
		c.Data["json"] = lists
		c.ServeJSON()
		c.StopRun()
	}else {
		message.Query().Filter("Open",1).Limit(4,0).All(&messages)
		lists = c.getMessag(messages)
	}

	//留言相册
	album := new(models.Album)
	var albums []*models.Album
	album.Query().Limit(6,0).OrderBy("-Views").All(&albums)
	//文章分类总数
	ArticleTagsCount := make(map[string]int64)
	Article := new(models.Article)
	ArticleTagsCount["study"],_ = Article.Query().Filter("Tags","study").Filter("Status",2).Count()
	ArticleTagsCount["journal"],_ = Article.Query().Filter("Tags","journal").Filter("Status",2).Count()
	ArticleTagsCount["life"],_ = Article.Query().Filter("Tags","life").Filter("Status",2).Count()
	ArticleTagsCount["bellesLettres"],_ = Article.Query().Filter("Tags","bellesLettres").Filter("Status",2).Count()

	//留言总数
	messageCount,_ :=  message.Query().Count()
	c.Data["albums"] = albums
	c.Data["ArticleTagsCount"] = ArticleTagsCount
	c.Data["messages"] = lists
	c.Data["dataTotal"] = messageCount
	c.Data["pageTotal"] = math.Ceil(float64(messageCount)/4)
	c.display("home/gbook")
}

func (c *GbookController) Add() {
	if c.Ctx.Request.Method =="POST"{
		name := c.GetString("name")
		email := c.GetString("email")
		content := c.GetString("lytext")
		if name == ""{
			c.ajaxMsg("请留下你的大名",501)
		}
		if email == ""{
			c.ajaxMsg("请留下你的邮箱",502)
		}
		if content == ""{
			c.ajaxMsg("请留下你想说的话",504)
		}
		message := new(models.Messages)
		message.Name = name
		message.Email = email
		message.Content = content
		message.Open = 1
		message.CreateTime = time.Now().UTC()
		if err := message.Insert(); err != nil{
			c.ajaxMsg(err.Error(),500)
		}
		c.ajaxMsg("留言成功",200)
	}
}

func (c *GbookController) getMessag(messages []*models.Messages) []*list {
	var lists []*list
	for _,v :=range messages{
		var (
			name string
			avatar string
		)
		if v.User != nil{
			user := new(models.User)
			user.Id = v.User.Id
			user.Read()
			name = user.UserName
			avatar = user.Avatar
		}else {
			name = v.Name
			avatar = "http://172.16.70.228:8080/group1/default/20190726/22/07/5/001QuqBczy7mIgPyRT248&690.jpg"
		}
		lisrobj := new(list)
		lisrobj.Name = name
		lisrobj.Avatar = avatar
		lisrobj.CreateTime = v.CreateTime.Format("2006-01-02")
		lisrobj.Content = v.Content
		lisrobj.Reply = v.Reply
		lists = append(lists, lisrobj)
	}
	return lists
}
