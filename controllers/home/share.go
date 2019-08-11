package home

import (
	"math"
	"simple_blog/models"
)

type ShareController struct {
	BaseController
}

func (c *ShareController) Index()  {
	album := new(models.Album)
	var albumList []*models.Album
	if c.Ctx.Request.Method == "POST"{
		page,_ := c.GetInt("page")
		album.Query().Filter("Status",2).Limit(8,(page-1)*8).All(&albumList)
		c.Data["json"] = albumList
		c.ServeJSON()
	}else {
		album.Query().Filter("Status",2).Limit(8,0).All(&albumList)
		albumCount,_ := album.Query().Filter("Status",2).Count()
		c.Data["albumList"] = albumList
		c.Data["pageTotal"] = math.Ceil(float64(albumCount)/8)
		c.Data["dataTotal"] = albumCount
		c.display("home/share")
	}

}