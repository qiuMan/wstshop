package models

import (
	"fmt"
	// "reflect"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BaseQueryParam struct {
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Offset int  `json:"offset"`
	Limit  int    `json:"limit"`
}

type Base struct {
	
}

func init() {
	fmt.Println("models/init")
	orm.RegisterModel(new(Users), new(Privileges), new(Menus))
}

func GetTableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix") 
	return prefix + name
}

func Db(Table string) (orm.QuerySeter) {
	query := orm.NewOrm().QueryTable(Table)

	return query
}

func GetQueryBuilder() (orm.QueryBuilder){
	db_type := beego.AppConfig.String("db_type")
	qb, _ := orm.NewQueryBuilder(db_type)
	return qb	
}

func Echo(values ...interface{}) {
	fmt.Println(values)
}

