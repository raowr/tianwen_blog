package admin

import (
	"tianwen_blog/models"
)

type AlbumController struct {
	BaseController
}

func (c *AlbumController) List()  {
	page,pagesize := c.getPage()
	var albums []*models.Album
	album := new(models.Album)
	album.Query().Limit(pagesize,(page-1)*pagesize).All(&albums)
	if c.Ctx.Request.Method == "POST"{
		c.Data["json"] = albums
		c.ServeJSON()
	}
	c.Data["list"] = albums
	c.Data["total"],_ = album.Query().Count()
	c.display("admin/album/list")
}

func (c *AlbumController) Add()  {
	if c.Ctx.Request.Method == "POST"{
		title := c.GetString("title")
		status,_ := c.GetInt("status")
		describe := c.GetString("content")
		cover := c.GetString("cover")
		album := new(models.Album)
		album.Title = title
		album.Cover = cover
		album.Status = status
		album.Describe = describe
		album.Updated = c.getTime()
		album.CreateTime = c.getTime()
		re := make(map[string]interface{})
		re["code"] = 200
		re["msg"] = "添加成功"
		if err := album.Inster(); err!=nil{
			re["code"] = 201
			re["msg"] = "添加失败"
		}
		c.Data["json"] = re
		c.ServeJSON()
	}
	c.display("admin/album/add")
	//c.StopRun()
	//var obj interface{}
	//req:=httplib.Post("http://172.16.70.214:8080/upload")
	//req.PostFile("file","./static/img/1.jpg")//注意不是全路径
	//req.Param("output","json")
	//req.Param("scene","")
	//req.Param("path","")
	//req.ToJSON(&obj)
	//fmt.Print(obj)
}

func (c *AlbumController) Edit()  {
	id,_:= c.GetInt("id")
	album := models.Album{Id:id}
	if c.Ctx.Request.Method == "POST"{
		albumData := models.Album{}
		album.Query().One(&albumData,"CreateTime")
		album.Title = c.GetString("title")
		album.Status,_ = c.GetInt("status")
		album.Describe = c.GetString("content")
		album.Cover = c.GetString("cover")
		album.Updated = c.getTime()
		album.CreateTime = albumData.CreateTime
		re := make(map[string]interface{})
		re["code"] = 200
		re["msg"] = "编辑成功"
		if err := album.Update(); err !=nil{
			re["code"] = 201
			re["msg"] = err
		}
		c.Data["json"] = re
		c.ServeJSON()
	}
	album.Read()
	c.Data["album"] = album
	c.display("admin/album/edit")
}