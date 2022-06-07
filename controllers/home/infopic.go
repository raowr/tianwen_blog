package home

import (
	"math"
	"tianwen_blog/models"
	"time"
)

type InfopicController struct {
	BaseController
}

func (c *InfopicController) Index() {
	id, _ := c.GetInt("id")
	//留言相册
	album := new(models.Album)
	var albums []*models.Album
	album.Query().Limit(6, 0).OrderBy("-Views").All(&albums)
	//文章分类总数
	ArticleTagsCount := make(map[string]int64)
	Article := new(models.Article)
	ArticleTagsCount["study"], _ = Article.Query().Filter("Tags", "study").Filter("Status", 2).Count()
	ArticleTagsCount["journal"], _ = Article.Query().Filter("Tags", "journal").Filter("Status", 2).Count()
	ArticleTagsCount["life"], _ = Article.Query().Filter("Tags", "life").Filter("Status", 2).Count()
	ArticleTagsCount["bellesLettres"], _ = Article.Query().Filter("Tags", "bellesLettres").Filter("Status", 2).Count()
	//id,_ := c.GetInt("id",0)
	albumInfo := make(map[string]interface{})
	album.Id = id
	if err := album.Read(); err != nil {
		c.Redirect("/error", 302)
		c.StopRun()
	}
	//增加浏览数
	album.Views++
	album.Update("Views")
	//相册内容
	albumInfo["Id"] = album.Id
	albumInfo["Title"] = album.Title
	var albumCreaTime string
	albumCreaTime = album.CreateTime.Format("2006-01-02")
	//photo列表
	photo := new(models.Photo)
	var photos []*models.Photo
	photo.Query().Filter("album__id", id).All(&photos)
	//上一篇
	album.Query().Filter("Id__lt", id).Filter("Status", 2).Limit(1, 0).OrderBy("-Id").One(album, "Id", "Title")
	albumInfo["prefixId"] = album.Id
	albumInfo["prefixTitle"] = album.Title
	//下一篇
	album.Query().Filter("Id__gt", id).Filter("Status", 2).Limit(1, 0).One(album, "Id", "Title")
	albumInfo["nextId"] = album.Id
	albumInfo["nextTitle"] = album.Title
	//总相片数
	photosCount, _ := photo.Query().Filter("album__id", id).Count()
	//评论列表
	comment := new(models.Comment)
	var comments []*models.Comment
	comment.Query().Filter("FirstId", id).Filter("Type", 2).Limit(6, 0).All(&comments)
	commentList := make(map[int]interface{})
	for i, v := range comments {
		commentMap := make(map[string]string)
		commentMap["Name"] = v.Name
		commentMap["Content"] = v.Content
		commentMap["CreateTime"] = v.CreateTime.Format("2006-01-02")
		commentList[i] = commentMap
	}
	commentCount, _ := comment.Query().Filter("FirstId", id).Filter("Type", 2).Count()

	c.Data["album"] = album
	c.Data["albumCreaTime"] = albumCreaTime
	c.Data["albums"] = albums
	c.Data["photos"] = photos
	c.Data["photosCount"] = photosCount
	c.Data["pageTotal"] = math.Ceil(float64(photosCount) / 6)
	c.Data["commentCount"] = commentCount
	c.Data["commentPageTotal"] = math.Ceil(float64(commentCount) / 6)
	c.Data["ArticleTagsCount"] = ArticleTagsCount
	c.Data["albumInfo"] = albumInfo
	c.Data["commentList"] = commentList
	c.display("home/infopic")
}

func (c *InfopicController) Photos() {
	id, _ := c.GetInt("id")
	page, _ := c.GetInt("page")
	//photo列表
	photo := new(models.Photo)
	var photos []*models.Photo
	photo.Query().Filter("album__id", id).Limit(6, (page-1)*6).All(&photos)
	c.Data["json"] = photos
	c.ServeJSON()
}

func (c *InfopicController) Comment() {
	id, _ := c.GetInt("id")
	page, _ := c.GetInt("page")
	//评论列表
	comment := new(models.Comment)
	var comments []*models.Comment
	comment.Query().Filter("FirstId", id).Filter("Type", 2).Limit(6, (page-1)*6).All(&comments)
	commentList := make(map[int]interface{})
	for i, v := range comments {
		commentMap := make(map[string]string)
		commentMap["Name"] = v.Name
		commentMap["Content"] = v.Content
		commentMap["CreateTime"] = v.CreateTime.Format("2006-01-02")
		commentList[i] = commentMap
	}
	c.Data["json"] = commentList
	c.ServeJSON()
}

func (c *InfopicController) CommentAdd() {
	id, _ := c.GetInt("id")
	username := c.GetString("name")
	saytext := c.GetString("saytext")
	if id == 0 {
		c.ajaxMsg("相册不存在", 401)
	}
	if username == "" {
		c.ajaxMsg("记得填姓名哦", 402)
	}
	if saytext == "" {
		c.ajaxMsg("没有想说的吗", 404)
	}
	comment := new(models.Comment)
	comment.FirstId = id
	comment.Type = 2
	comment.ParentId = 0
	comment.Name = username
	comment.Content = saytext
	comment.CreateTime = time.Now().UTC()
	if err := comment.Insert(); err != nil {
		c.ajaxMsg(err.Error(), 400)
	}
	c.ajaxMsg("评论成功", 200)

}
