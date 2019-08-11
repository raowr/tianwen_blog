package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id				int
	User			*User		`orm:"rel(fk);null;description(评论会员)"`
	Name 			string		`orm:"size(255);null;description(评论人姓名)"`
	FirstId			int			`orm:"description(顶级id)"`
	Type			int			`orm:"description(评论类型1=文章 2=相册)"`
	ParentId		int			`orm:"description(父级id)"`
	Content			string		`orm:"type(text);description(评论内容)"`
	CreateTime		time.Time	`orm:"type(datetime);description(评论时间)"`
}

func (m *Comment) GetTableName() string {
	return TableName("Comment")
}

func (m *Comment) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil{
		return err
	}
	return nil
}

func (m *Comment) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err != nil{
		return err
	}
	return nil
}

func (m *Comment) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...);err != nil {
		return err
	}
	return nil
}

func (m *Comment)Delete() error  {
	if _,err := orm.NewOrm().Delete(m); err != nil{
		return err
	}
	return nil
}

func (m *Comment) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}


