package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id                 int
	RealName           string `orm:"size(32);null;description(真实姓名)" `
	UserName           string `orm:"size(24);description(用户名)"`
	UserPwd            string `orm:"size(255);description(密码)"`
	IsSuper            bool	  `orm:"null;description(是否超级会员)" `
	Status             int	  `orm:"default(1);description(状态)" `
	Mobile             string `orm:"size(16);null;description(手机号码)" `
	Email              string `orm:"size(256);null; description(电子邮件)"`
	Avatar             string `orm:"size(256);null;description(头像)"`
	Sex				   string `orm:"size(10);null; description(性别)"`
	Birthday		   time.Time `orm:"type(date);null;description(生日)"`
	Messages		   []*Messages	`orm:"reverse(many)"`		//多个留言
}


func (m *User) GetTableName() string {
	return TableName("user")
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil{
		return err
	}
	return nil
}

func (m *User) Insert() (int64,error) {
	var (
		id int64
		err error
	)
	if id,err = orm.NewOrm().Insert(m); err != nil {
		return 0,err
	}
	return id,nil
}

func (m *User) Update(cols ...string) (int64,error) {
	var id int64
	var err error
	if id,err = orm.NewOrm().Update(m,cols...);err != nil{
		return 0,err
	}
	return id,nil
}

func (m *User) Delete() (int64,error){
	var id int64
	var err error
	if id,err = orm.NewOrm().Delete(m); err != nil{
		return 0,err
	}
	return id,nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

