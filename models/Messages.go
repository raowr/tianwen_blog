package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Messages struct {
	Id			int
	User		*User			`orm:"rel(fk);null;description(留言会员)"`
	Name		string			`orm:"size(255);null;description(名称)"`
	Email		string			`orm:"size(255);null;description(邮箱)"`
	Wx			string			`orm:"size(255);null;description(微信)"`
	Content		string			`orm:"type(longtext);null;description(留言内容)"`
	Open		int				`orm:"default(1);description(留言审核:1=审核 0=不审核)"`
	Reply		string			`orm:"type(longtext);null;description(站长回复)"`
	CreateTime	time.Time		`orm:"type(datetime);null;description(留言时间)"`
}

func (m *Messages) GetTableName() string {
	return TableName("Messages")
}

func (m *Messages) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil{
		return err
	}
	return nil
}

func (m *Messages) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err != nil{
		return err
	}
	return nil
}

func (m *Messages) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...);err != nil {
		return err
	}
	return nil
}

func (m *Messages)Delete() error  {
	if _,err := orm.NewOrm().Delete(m); err != nil{
		return err
	}
	return nil
}

func (m *Messages) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

