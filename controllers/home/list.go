package home

import (
	"math"
	"simple_blog/models"
)

type ListController struct {
	BaseController
}

func (c *ListController) Index() {
	Article := new(models.Article)
	var articles []*models.Article
	if c.Ctx.Request.Method == "POST" {
		//文章列表 post
		page, _ := c.GetInt("page")
		Article.Query().Filter("Status", 2).OrderBy("-Views").Limit(10, (page-1)*10).All(&articles)
		for i, v := range articles {
			articles[i].Content = string([]rune(v.Content)[:80])
		}
		c.Data["json"] = articles
		c.ServeJSON()
		c.StopRun()
	}
	//文章分类总数
	ArticleTagsCount := make(map[string]int64)
	ArticleTagsCount["study"], _ = Article.Query().Filter("Tags", "study").Filter("Status", 2).Count()
	ArticleTagsCount["journal"], _ = Article.Query().Filter("Tags", "journal").Filter("Status", 2).Count()
	ArticleTagsCount["life"], _ = Article.Query().Filter("Tags", "life").Filter("Status", 2).Count()
	ArticleTagsCount["bellesLettres"], _ = Article.Query().Filter("Tags", "bellesLettres").Filter("Status", 2).Count()
	//站长推荐
	var ArticleRecommend []*models.Article
	Article.Query().Filter("Recommend", 1).Filter("Status", 2).Limit(8, 0).All(&ArticleRecommend)
	//点击排行榜
	var ArticleViews []*models.Article
	Article.Query().OrderBy("-Views").Filter("Status", 2).Limit(8, 0).All(&ArticleViews)
	//文章列表
	Article.Query().Filter("Status", 2).OrderBy("-Views").Limit(10, 0).All(&articles)
	for i, v := range articles {
		articles[i].Content = string([]rune(v.Content)[:80])
	}
	articleCount, _ := Article.Query().Filter("Status", 2).Count()
	c.Data["ArticleTagsCount"] = ArticleTagsCount
	c.Data["ArticleRecommend"] = ArticleRecommend
	c.Data["ArticleViews"] = ArticleViews
	c.Data["articles"] = articles
	c.Data["pageTotal"] = math.Ceil(float64(articleCount) / 10)
	c.Data["dataTotal"] = articleCount
	c.display("home/list")
}
