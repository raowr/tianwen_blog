package home

import (
	"tianwen_blog/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	//首页相册
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
	//站长推荐
	var ArticleRecommend []*models.Article
	Article.Query().Filter("Recommend", 1).Filter("Status", 2).Limit(8, 0).All(&ArticleRecommend)
	//文章列表
	var articles []*models.Article
	Article.Query().Filter("Status", 2).OrderBy("-Views").Limit(10, 0).All(&articles)
	c.Data["albums"] = albums
	c.Data["ArticleTagsCount"] = ArticleTagsCount
	c.Data["ArticleRecommend"] = ArticleRecommend
	c.Data["articles"] = articles
	c.display("home/index")
}
