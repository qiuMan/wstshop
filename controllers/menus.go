package controllers

import "wstmart/models"

type Menus struct {
	Base
}

func (this *Menus) Index() {
	if this.Ctx.Input.IsPost() {
		M := new(models.Menus)
		data := M.ListQuery()
		this.JsonOut(data, len(data))
	} 
}