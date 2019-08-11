package models

import "github.com/astaxie/beego/orm"

type Photo struct {
	Id		int
	Url		string		`orm:"size(255)"`
	Data	string		`orm:"type(text)"`
	Album	*Album		`orm:"rel(fk);null"`
}

func (m *Photo) GetTableName() string {
	return TableName("photo")
}

func (m *Photo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil{
		return err
	}
	return nil
}

func (m *Photo) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err != nil{
		return err
	}
	return nil
}

func (m *Photo) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...);err != nil {
		return err
	}
	return nil
}

func (m *Photo)Delete() error  {
	if _,err := orm.NewOrm().Delete(m); err != nil{
		return err
	}
	return nil
}

func (m *Photo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

