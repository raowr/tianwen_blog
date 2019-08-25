package admin

import (
	"simple_blog/models"
	"strconv"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) List() {
	page, pagesize := c.getPage()
	var articles []*models.Article
	var total int64
	article := new(models.Article)
	articleObj := article.Query().Limit(pagesize, (page-1)*pagesize)
	if c.Ctx.Request.Method == "POST" {
		list := make(map[string]interface{})
		keyword := c.GetString("keyword")
		status, _ := c.GetInt("statusValue", 0)
		articleTotal := article.Query()
		if keyword != "" {
			articleObj = articleObj.Filter("title__icontains", keyword)
			articleTotal = article.Query().Filter("title__icontains", keyword)
		}
		if status != 0 {
			articleObj = articleObj.Filter("Status", status)
			articleTotal = article.Query().Filter("Status", status)
		}
		articleObj.All(&articles)
		total, _ = articleTotal.Count()
		for i, value := range articles {
			if len(value.Briefly) > 60 {
				articles[i].Briefly = string([]rune(value.Briefly)[:60])
			}
		}

		list["articles"] = articles
		list["total"] = total
		c.Data["json"] = list
		c.ServeJSON()
	} else {
		articleObj.All(&articles)
		for i, value := range articles {
			if len(value.Briefly) > 60 {
				articles[i].Briefly = string([]rune(value.Briefly)[:60])
			}
		}
		total, _ = article.Query().Count()
		c.Data["list"] = articles
		c.Data["total"] = total
		c.display("admin/article/list")
	}
}

func (c *ArticleController) Add() {
	if c.Ctx.Request.Method == "POST" {
		title := c.GetString("title")
		tags := c.GetString("tags")
		status, _ := c.GetInt("status")
		briefly := c.GetString("briefly")
		content := c.GetString("content")
		bannerUrl := c.GetString("bannerUrl")
		recommend, _ := c.GetInt("recommend")
		adminInfo := c.GetSession("admin")
		admin, _ := adminInfo.(models.Admin)

		Article := new(models.Article)
		Article.Title = title
		Article.Tags = tags
		Article.Status = status
		Article.Briefly = briefly
		Article.Content = content
		Article.BannerUrl = bannerUrl
		Article.Recommend = recommend
		Article.Author = strconv.Itoa(admin.Id)
		Article.Updated = c.getTime()
		Article.CreateTime = c.getTime()
		re := make(map[string]interface{})
		re["code"] = 200
		re["msg"] = "发布成功"
		if err := Article.Insert(); err != nil {
			re["code"] = 201
			re["msg"] = err
		}
		c.Data["json"] = re
		c.ServeJSON()

	}
	c.display("admin/article/add")
}

func (c *ArticleController) Edit() {
	id, _ := strconv.Atoi(c.GetString("id"))
	article := models.Article{Id: id}
	articleData := models.Article{}
	if c.Ctx.Request.Method == "POST" {
		article.Query().One(&articleData, "CreateTime")
		title := c.GetString("title")
		tags := c.GetString("tags")
		status, _ := c.GetInt("status")
		briefly := c.GetString("briefly")
		content := c.GetString("content")
		bannerUrl := c.GetString("bannerUrl")
		recommend, _ := c.GetInt("recommend")
		article.Title = title
		article.Tags = tags
		article.Status = status
		article.Briefly = briefly
		article.Content = content
		article.BannerUrl = bannerUrl
		article.Recommend = recommend
		article.Updated = c.getTime()
		article.CreateTime = articleData.CreateTime
		re := make(map[string]interface{})
		re["code"] = 200
		re["msg"] = "编辑成功"
		if err := article.Update(); err != nil {
			re["code"] = 201
			re["msg"] = err
		}
		c.Data["json"] = re
		c.ServeJSON()
	}
	article.Read()
	c.Data["article"] = article
	c.display("admin/article/edit")
}
