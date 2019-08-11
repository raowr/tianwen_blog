package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init()  {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")

	if dbhost == ""{
		dbhost = "3306"
	}
	orm.DefaultTimeLoc = time.UTC
	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Local"
	orm.RegisterDataBase("default", "mysql", dburl)
	orm.RegisterModel(
		new(Article),
		new(Album),
		new(Photo),
		new(User),
		new(Messages),
		new(Admin),
		new(Comment))
	orm.RunSyncdb("default", false, true)
}
