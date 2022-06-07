package home

import "tianwen_blog/models"

type AboutController struct {
	BaseController
}

func (c *AboutController) Index()  {
	//关于我相册
	album := new(models.Album)
	var albums []*models.Album
	album.Query().Limit(6,0).OrderBy("-Views").All(&albums)
	//分类，文章
	Article := new(models.Article)
	c.Data["albums"] = albums
	c.Data["articleTotal"],_ = Article.Query().Filter("Status",2).Count()
	c.Data["albumTotal"],_ = album.Query().Filter("Status",2).Count()
	c.display("home/about")
}
