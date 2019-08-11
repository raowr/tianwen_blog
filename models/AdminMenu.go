package models

import (
	"fmt"
	"github.com/astaxie/beego"
)

type AdminMenu struct {
	ID            	int64        	`orm:"column(id);pk;auto"`
	ParentId		int64			`orm:"default(1)"`
	Title			string			`orm:"null"`
	Url				string			`orm:"null"`
	IsUse			bool			`orm:"default(false)"`
}

func (this *AdminMenu) TableName() string {
	return TableName("AdminMenu")
}

//返回带前缀的表名
func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}


