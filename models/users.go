package models

import (
	// "fmt"
	// "reflect"
	"github.com/astaxie/beego/orm"
)

type UsersQueryParam struct {
	BaseQueryParam
	UserNameLike string //模糊查询
	PhoneLike string //模糊查询
}

type Users struct {
	Id 				int 		`orm:"pk;column(userId);"`
	LoginName 		string 		`orm:"column(loginName)"`
	LoginPwd 		string 		`orm:"column(loginPwd)"`
	loginSecret 	int 		`orm:"column(LoginSecret)"`
	UserSex 		int8 		`orm:"column(userSex)"`
	UserType 		int8 		`orm:"column(userType)"`
	UserName 		string 		`orm:"column(userName)"`
	TrueName 		string 		`orm:"column(trueName)"`
	UserPhone 		string  	`orm:"column(userPhone)"`
	Brithday 		string  	`orm:"column(brithday)"`
	UserPhoto 		string  	`orm:"column(userPhoto)"`
	UserQQ 			string  	`orm:"column(userQQ)"`
	UserEmail 		string  	`orm:"column(userEmail)"`
	UserScore 		string  	`orm:"column(userScore)"`
	UserTotalScore 	string  	`orm:"column(userTotalScore)"`
	LastIP 			string  	`orm:"column(lastIP)"`
	LastTime 		string  	`orm:"column(lastTime)"`
	UserFrom 		int8  		`orm:"column(userFrom)"`
	UserMoney 		float64  	`orm:"column(userMoney)"`
	LockMoney 		float64  	`orm:"column(lockMoney)"`
	UserStatus 		int8	  	`orm:"column(userStatus)"`
	DataFlag 		int8	  	`orm:"column(dataFlag)"`
	CreateTime 		string  	`orm:"column(createTime)"`
	PayPwd 			string  	`orm:"column(payPwd)"`
	Table 			string  	`orm:"column(table)"`
}

func NewUsers() *Users {
	obj := new(Users)
	return obj 
}

func (this *Users) TableName() string {
	// db table name
    return GetTableName("users")
}

func (this *Users) PageQuery(Params *UsersQueryParam) ([]*Users, int64) {
	table := this.TableName()
	query := Db(table)

	if Params.UserNameLike != "" {
		query = query.Filter("loginName__icontains", Params.UserNameLike)
	}

	if Params.PhoneLike != "" {
		query = query.Filter("userPhone__icontains", Params.PhoneLike)
	}

	total, _ := query.Count()
	data := make([]*Users, 0)
	
	if total > 0 {
		query.Limit(Params.Limit, Params.Offset).All(&data)
	}

	return data, total
	
} 

func (this *Users) GetUsetById(id int) (*Users, error) {
	m := Users{Id: id}
	o := orm.NewOrm()
	err := o.Read(&m)
	
	if err != nil {
		return nil, err
	}

	return &m, err

}

func (this *Users) GetUserInfoByName(username string) (Users, error) {
	var user Users
	table := this.TableName()
	
	o := orm.NewOrm()
	err := o.QueryTable(table).Filter("loginName", username).One(&user)
	
	return user, err
	
}

