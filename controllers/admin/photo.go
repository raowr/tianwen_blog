package admin

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"log"
	"simple_blog/models"
	"strconv"
	"strings"
)

type PhotoController struct {
	BaseController
}

func (c *PhotoController) Upload()  {
	album := new(models.Album)
	albumId ,_ := c.GetInt("title")
	album.Id=albumId
	f, h, err := c.GetFile("photo")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	file := "static/upload/" + h.Filename
	defer f.Close()
	c.SaveToFile("photo", file) // 保存位置在 static/upload, 没有文件夹要先创建
	var obj map[string]string
	var photo models.Photo
	uploadUrl := beego.AppConfig.String("uploadUrl")
	req:=httplib.Post(uploadUrl)
	req.PostFile("file",file)//注意是全路径
	req.Param("output","json")
	req.Param("scene","")
	req.Param("path","")
	req.ToJSON(&obj)
	photo.Url=obj["url"]
	data,_ := json.Marshal(obj)
	photo.Data = string(data)
	photo.Album = album
	re := make(map[string]interface{})
	re["code"] = 200
	re["msg"] = "添加成功"
	if err := photo.Insert(); err!=nil{
		re["code"] = 201
		re["msg"] = "添加失败"
		re["error"] = err.Error()
	}else {
		re["id"] = photo.Id
	}
	c.Data["json"] = re
	c.ServeJSON()
}

func (c *PhotoController) List()  {
	photo := new(models.Photo)
	albun := new(models.Album)
	var photoList  []*models.Photo
	albumId ,_ := c.GetInt("id")
	albun.Id = albumId
	photo.Album = albun
	photo.Query().Filter("Album",albun).All(&photoList)
	for _,v := range photoList{
		fmt.Println(v)
	}
	c.Data["list"] = photoList
	c.display("admin/photo/list")
}

func (c *PhotoController) Add()  {
	c.display("admin/photo/add")
}

func (c *PhotoController) Update()  {
	photo := new(models.Photo)
	album := new(models.Album)
	albumId ,_ := c.GetInt("title")
	album.Id = albumId
	ids := c.GetString("ids")
	idslice := strings.Split(ids, ",")
	photo.Album = album
	for  _,value := range idslice{
		photo.Id,_ = strconv.Atoi(value)
		photo.Update("Album")
	}
	re := make(map[string]interface{})
	re["code"] = 200
	re["msg"] = "添加成功"
	c.Data["json"] = re
	c.ServeJSON()
}

func (c *PhotoController) Delete()  {
	Id ,_ := c.GetInt("id")
	photo := models.Photo{Id:Id}
	re := make(map[string]interface{})
	re["code"] = 200
	re["msg"] = "删除成功"
	if err := photo.Delete();err != nil{
		re["code"] = 201
		re["msg"] = "删除失败"
		re["error"] = err.Error()
	}
	c.Data["json"] = re
	c.ServeJSON()
}

func (c *PhotoController) Cancel()  {
	photo := new(models.Photo)
	ids := c.GetString("ids")
	idslice := strings.Split(ids, ",")
	re := make(map[string]interface{})
	for  _,value := range idslice{
		photo.Id,_ = strconv.Atoi(value)
		photo.Delete()
	}
	re["code"] = 200
	re["msg"] = "取消完成"
	c.Data["json"] = re
	c.ServeJSON()
}
