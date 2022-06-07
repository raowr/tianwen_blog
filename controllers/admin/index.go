package admin

import "tianwen_blog/models"

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	adminInfo := c.GetSession("admin")
	admin, _ := adminInfo.(models.Admin)
	article := new(models.Article)
	user := new(models.User)
	comment := new(models.Comment)
	photo := new(models.Photo)
	c.Data["admin"] = admin
	c.Data["articleNum"],_ = article.Query().Count()
	c.Data["userNum"],_ = user.Query().Count()
	c.Data["commentNum"],_ = comment.Query().Count()
	c.Data["photoNum"],_ = photo.Query().Count()
	c.display("")
}
