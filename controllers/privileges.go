/*
 * @Author: v_wenlqiu 
 * @Date: 2019-02-02 13:09:14 
 * @Last Modified by: v_wenlqiu
 * @Last Modified time: 2019-02-03 21:37:47
 */
package controllers

import "wstmart/models"

type Privileges struct {
	Base
}

func (this *Privileges) Index()  {
	parentId, _ := this.GetInt(":id", 0)
	menus := new(models.Menus)
	data := menus.GetById(parentId)

	if this.Ctx.Input.IsPost() {
		m := new(models.Privileges)
		data := m.ListQuery(parentId)
		this.JsonOut(data, len(data))
	}

	this.Data["data"] = data 
}

func (this *Privileges) Get() {
	id, _ := this.GetInt(":id", 0)
	m := new(models.Privileges)
	data := m.GetById(id)
	Echo(data)
	this.Data["data"] = data
	this.TplName = "privileges/detail.html"
	this.Layout = "layout/open.html"
}

func (this *Privileges) Post() {
	mid, _ := this.GetInt(":id", 0)
	m := models.NewPrivileges()
	m.MenuId = mid
	m.PrivilegeCode = this.GetString("code")
	m.PrivilegeName = this.GetString("name")
	m.IsMenuPrivilege, _ = this.GetInt8("is-menu", 0)
	m.PrivilegeUrl = this.GetString("url")
	m.OtherPrivilegeUrl = this.GetString("other-url")
	m.DataFlag = 1
	
	id,err := m.Add(m)
	code := 0
	if err != nil {
		code = 99
	}
	this.JsonOut(m, id, code, err)
}

func (this *Privileges) Put() {
	id, _ := this.GetInt(":id", 0)
	filters := make(map[string]interface{})
	filters["privilegeId"] = id
	data := make(map[string]interface{})
	data["privilegeCode"] = this.GetString("code")
	data["privilegeName"] = this.GetString("name")
	data["isMenuPrivilege"], _ = this.GetInt("is-menu", 0)
	data["privilegeUrl"] = this.GetString("url")
	data["otherPrivilegeUrl"] = this.GetString("other-url")

	m := models.NewPrivileges()
	num, err := m.Update(filters, data)
	code := 0
	if err != nil {
		code = 99
	} 
	this.JsonOut(data, num, code, err)
}

func (this *Privileges) Delete () {
	id, _ := this.GetInt(":id", 0)
	m := models.NewPrivileges()
	num, err := m.Delete(id)
	code := 0 
	
	if err != nil {
		code = 99
	}

	this.JsonOut(id, num, code, err)

}