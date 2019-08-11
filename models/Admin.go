package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Admin struct {
	Id			int
	UserName 	string		`orm:"description(登录名)"`
	Pwd			string		`orm:"description(登录密码)"`
	GroupId		int			`orm:"null;default(1);description(分组id)"`
	Email		string		`orm:"null;description(邮箱)"`
	RealName 	string		`orm:"null;description(真实姓名)"`
	Tel			string		`orm:"null;description(联系号码)"`
	CreateTime	time.Time	`orm:"type(datetime);null;description(添加时间)"`
	Status		int			`orm:"default(1);null;description(状态)"`
	Avatar		string		`orm:"null;description(头像)"`
}


func (m *Admin) GetTableName() string {
	return TableName("Admin")
}

func (m *Admin) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil{
		return err
	}
	return nil
}

func (m *Admin) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err != nil{
		return err
	}
	return nil
}

func (m *Admin) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...);err != nil {
		return err
	}
	return nil
}

func (m *Admin)Delete() error  {
	if _,err := orm.NewOrm().Delete(m); err != nil{
		return err
	}
	return nil
}

func (m *Admin) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

