package sysinit

import (	
	"fmt"
	_ "wstmart/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//配置加载
	setConfig()
	//数据库连接
	dbConn()
	
}

func setConfig() {
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	
}

func dbConn() {
	//数据库名称
	dbName := beego.AppConfig.String("db_name")
	fmt.Println(dbName)
	//数据库连接用户名
	dbUser := beego.AppConfig.String("db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String("db_pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String("db_host")
	//数据库端口
	dbPort := beego.AppConfig.String("db_port")

	//注册mysql Driver
    // orm.RegisterDriver("mysql", orm.DR_MySQL)
	// orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.224.128:3306)/cloudta?charset=utf8")
    conn := dbUser + ":" +dbPwd + "@tcp" + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8" 
    //连接
    orm.RegisterDataBase("default", "mysql", conn)

	orm.RunSyncdb("default", false, true)
}