package controllers

import (
	// "fmt"
	"wstmart/models"
)


type Users struct {
	Base
}

func (this *Users) Index() {
	if this.Ctx.Input.IsPost() {
		queryParam := new(models.UsersQueryParam)
		queryParam.Limit, _ = this.GetInt("limit", 10)
		page, _ := this.GetInt("page", 1)
		queryParam.Offset = (page - 1) * queryParam.Limit
		queryParam.UserNameLike = this.GetString("name")
		queryParam.PhoneLike = this.GetString("phone")
		
		M := models.NewUsers()
		
		data, count := M.PageQuery(queryParam)
		
		this.JsonOut(data, count)
		
	} 
}

func (this *Users) Get() {
	this.TplName = "users/detail.html"
	this.SetTpl()
}

func (this *Users) Post() {
	
}


