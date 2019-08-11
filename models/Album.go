package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Album struct {
	Id			int
	Title		string		`orm:"size(255)"`			//标题
	Describe	string		`orm:"type(text)"`			//描述
	Cover		string		`orm:"type(text)"`			//封面
	Views		int										//浏览数
	Status		int										//状态
	Updated  	time.Time 	`orm:"type(datetime)"`		//修改时间
	CreateTime	time.Time	`orm:"type(datetime)"`		//创建时间
	Photo		[]*Photo	`orm:"reverse(many)"`		//一个相册多张照片
}

func (m *Album) GetTableName() string {
	return TableName("album")
}

func (m *Album) Inster() error {
	if _,err := orm.NewOrm().Insert(m); err != nil{
		return err
	}
	return nil
}

func (m *Album) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil{
		return err
	}
	return nil
}

func (m *Album) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...); err != nil{
		return err
	}
	return nil
}

func (m *Album) Delete() error {
	if _,err := orm.NewOrm().Delete(m); err != nil{
		return err
	}
	return nil
}

func (m *Album) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
