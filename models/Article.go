package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id			int										//Id
	Author   	string 		`orm:"size(15)"`			//作者
	Title    	string 		`orm:"size(100)"`			//标题
	Content		string		`orm:"type(text)"`			//内容
	Tags		string		`orm:"size(100)"`			//分类
	BannerUrl	string		`orm:"type(text)"`			//右侧图
	Views		int										//浏览数
	Recommend	int										//是否推荐
	Status		int										//状态
	Updated  	time.Time 	`orm:"type(datetime)"`		//修改时间
	CreateTime	time.Time	`orm:"type(datetime)"`		//创建时间
}

func (m *Article)TableName() string {
	return TableName("article")
}

func (m *Article) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err != nil{
		return err
	}else {
		return nil
	}
}

func (m *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil{
		return err
	}else {
		return nil
	}
}

func (m *Article) Update (fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...); err != nil{
		return err
	}else {
		return nil
	}
}

func (m *Article) Delete() error {
	if _,err := orm.NewOrm().Delete(m);err != nil{
		return err
	}else {
		return nil
	}
}

func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}