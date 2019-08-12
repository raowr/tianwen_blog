package home

import (
	"math"
	"simple_blog/models"
)

type ShareController struct {
	BaseController
}

func (c *ShareController) Index() {
	album := new(models.Album)
	var albumList []*models.Album
	if c.Ctx.Request.Method == "POST" {
		page, _ := c.GetInt("page")
		album.Query().Filter("Status", 2).Limit(8, (page-1)*8).All(&albumList)
		albumlists := c.getAlbumlists(albumList)
		c.Data["json"] = albumlists
		c.ServeJSON()
	} else {
		album.Query().Filter("Status", 2).Limit(8, 0).All(&albumList)
		albumlists := c.getAlbumlists(albumList)
		albumCount, _ := album.Query().Filter("Status", 2).Count()
		c.Data["albumList"] = albumlists
		c.Data["pageTotal"] = math.Ceil(float64(albumCount) / 8)
		c.Data["dataTotal"] = albumCount
		c.display("home/share")
	}
}

func (c *ShareController) getAlbumlists(albumList []*models.Album) (albumlists []interface{}) {
	for i := 0; i < 4; i++ {
		albums := make(map[int]interface{})
		if i < len(albumList) {
			albums[0] = albumList[i]
		}
		if i+4 < len(albumList) {
			albums[1] = albumList[i+4]
		}
		albumlists = append(albumlists, albums)
	}
	return
}
