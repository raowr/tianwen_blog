package routers

import (
	"simple_blog/controllers/admin"
	"simple_blog/controllers/home"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Router("/", &home.IndexController{},"get:Index")
    beego.Router("/index", &home.IndexController{},"get:Index")
    beego.Router("/share", &home.ShareController{},"get,post:Index")
    beego.Router("/list", &home.ListController{},"get,post:Index")
    beego.Router("/about", &home.AboutController{},"get:Index")
    beego.Router("/gbook", &home.GbookController{},"get,post:Index")
    beego.Router("/gbook/add", &home.GbookController{},"get,post:Add")
    beego.Router("/info", &home.InfoController{},"get,post:Index")
    beego.Router("/info/comment_add", &home.InfoController{},"get,post:CommentAdd")
    beego.Router("/info/comment_get", &home.InfoController{},"get,post:CommentGet")
    beego.Router("/infopic", &home.InfopicController{},"get:Index")
    beego.Router("/infopic/photos", &home.InfopicController{},"post:Photos")
    beego.Router("/infopic/comment", &home.InfopicController{},"post:Comment")
    beego.Router("/infopic/comment_add", &home.InfopicController{},"post:CommentAdd")


    //admin路由
	beego.Router("/admin",&admin.IndexController{},"get:Index")
	beego.Router("/article/list",&admin.ArticleController{},"get,post:List")
	beego.Router("/article/add",&admin.ArticleController{},"get,post:Add")
	beego.Router("/article/edit",&admin.ArticleController{},"get,post:Edit")
	beego.Router("/album/list",&admin.AlbumController{},"get,post:List")
	beego.Router("/album/add",&admin.AlbumController{},"get,post:Add")
	beego.Router("/album/edit",&admin.AlbumController{},"get,post:Edit")
	beego.Router("/photo/upload",&admin.PhotoController{},"get,post:Upload")
	beego.Router("/photo/list",&admin.PhotoController{},"get,post:List")
	beego.Router("/photo/add",&admin.PhotoController{},"get,post:Add")
	beego.Router("/photo/update",&admin.PhotoController{},"post:Update")
	beego.Router("/photo/delete",&admin.PhotoController{},"post:Delete")
	beego.Router("/photo/cancel",&admin.PhotoController{},"post:Cancel")
	beego.Router("/user/list",&admin.UserController{},"get,post:List")
	beego.Router("/user/edit",&admin.UserController{},"get,post:Edit")
	beego.Router("/user/add",&admin.UserController{},"get,post:Add")
	beego.Router("/user/avatarUpload",&admin.UserController{},"get,post:AvatarUpload")
	beego.Router("/user/editPwd",&admin.UserController{},"get,post:EditPwd")
	beego.Router("/message/list",&admin.MessagesController{},"get,post:List")
	beego.Router("/comment/list",&admin.CommentController{},"get,post:List")
	beego.Router("/admin/login",&admin.LoginController{},"get,post:Login")
	beego.Router("/admin/loginOut",&admin.LoginController{},"get,post:LoginOut")
	beego.Router("/admin/profile",&admin.AdminController{},"get,post:Profile")
	beego.Router("/admin/settings",&admin.AdminController{},"get,post:Settings")
	beego.Router("/admin/editPwd",&admin.AdminController{},"get,post:EditPwd")
}
