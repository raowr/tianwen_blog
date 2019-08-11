package home

import (
	"math"
	"simple_blog/models"
	"time"
)

type InfoController struct {
	BaseController
}

type commentData struct {
	Name string
	Content string
	CreateTime string
}

func (c *InfoController) Index()  {
	Article := new(models.Article)
	//文章分类总数
	ArticleTagsCount := make(map[string]int64)
	ArticleTagsCount["study"],_ = Article.Query().Filter("Tags","study").Filter("Status",2).Count()
	ArticleTagsCount["journal"],_ = Article.Query().Filter("Tags","journal").Filter("Status",2).Count()
	ArticleTagsCount["life"],_ = Article.Query().Filter("Tags","life").Filter("Status",2).Count()
	ArticleTagsCount["bellesLettres"],_ = Article.Query().Filter("Tags","bellesLettres").Filter("Status",2).Count()
	//站长推荐
	var ArticleRecommend []*models.Article
	Article.Query().Filter("Recommend",1).Filter("Status",2).Limit(8,0).All(&ArticleRecommend)
	//点击排行榜
	var ArticleViews []*models.Article
	Article.Query().OrderBy("-Views").Filter("Status",2).Limit(8,0).All(&ArticleViews)
	//文章内容
	id,_ := c.GetInt("id")
	Article.Id = id
	Article.Read()
	ArticleInfo := make(map[string]interface{})
	ArticleInfo["id"] = Article.Id
	ArticleInfo["title"] = Article.Title
	ArticleInfo["author"] = Article.Author
	ArticleInfo["createTime"] = Article.CreateTime.Format("2006-01-02")
	var tags string
	if Article.Tags == "study" {
		tags = "学无止境"
	}else if Article.Tags == "life"{
		tags = "慢生活"
	}else if Article.Tags == "journal"{
		tags = "日记"
	}else if Article.Tags == "bellesLettres"{
		tags = "美文欣赏"
	}
	ArticleInfo["tags"] = tags
	ArticleInfo["views"] = Article.Views
	ArticleInfo["content"] = Article.Content
	//上一篇
	Article.Query().Filter("Id__lt",id).Filter("Status",2).Limit(1,0).OrderBy("-Id").One(Article,"Id","Title")
	ArticleInfo["prefixId"] = Article.Id
	ArticleInfo["prefixTitle"] = Article.Title
	//下一篇
	Article.Query().Filter("Id__gt",id).Filter("Status",2).Limit(1,0).One(Article,"Id","Title")
	ArticleInfo["nextId"] = Article.Id
	ArticleInfo["nextTitle"] = Article.Title
	//评论
	comment := new(models.Comment)
	var commentList []*models.Comment
	comment.Query().Filter("FirstId",id).Filter("Type",1).Limit(4,0).All(&commentList)
	var comments []*commentData
	for _,v := range commentList {
		commentDatas := new(commentData)
		commentDatas.Name = v.Name
		commentDatas.Content = v.Content
		commentDatas.CreateTime = v.CreateTime.Format("2006-01-02")
		comments = append(comments, commentDatas)
	}
	commentCount,_ :=  comment.Query().Filter("FirstId",id).Filter("Type",1).Count()
	c.Data["ArticleTagsCount"] = ArticleTagsCount
	c.Data["ArticleRecommend"] = ArticleRecommend
	c.Data["ArticleViews"] = ArticleViews
	c.Data["Article"] = ArticleInfo
	c.Data["comments"] = comments
	c.Data["pageTotal"] = math.Ceil(float64(commentCount)/4)
	c.Data["dataTotal"] = commentCount
	c.display("home/info")
}

func (c *InfoController) CommentAdd() {
	id,_ := c.GetInt("id")
	name := c.GetString("name")
	content := c.GetString("saytext")
	comment := new(models.Comment)
	comment.FirstId = id
	comment.Content = content
	comment.Name = name
	comment.Type = 1
	comment.ParentId = 0
	comment.CreateTime = time.Now().UTC()
	if err := comment.Insert();err != nil{
		c.ajaxMsg(err.Error(),200)
	}
	c.ajaxMsg("评论成功",200)
}

func (c *InfoController) CommentGet()  {
	id,_ := c.GetInt("id")
	page,_ := c.GetInt("page")
	comment := new(models.Comment)
	var commentList []*models.Comment
	comment.Query().Filter("FirstId",id).Filter("Type",1).Limit(4,(page-1)*4).All(&commentList)
	var comments []*commentData
	for _,v := range commentList {
		commentDatas := new(commentData)
		commentDatas.Name = v.Name
		commentDatas.Content = v.Content
		commentDatas.CreateTime = v.CreateTime.Format("2006-01-02")
		comments = append(comments, commentDatas)
	}
	c.Data["json"] = comments
	c.ServeJSON()
	c.StopRun()
}

