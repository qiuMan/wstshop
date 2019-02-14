/*
 * @Author: v_wenlqiu 
 * @Date: 2019-02-02 13:33:03 
 * @Last Modified by: v_wenlqiu
 * @Last Modified time: 2019-02-02 13:48:29
 */
package models

type Menus struct {
	MenuId		    int 	`orm:"pk;column(menuId);"`
	ParentId		int 	`orm:"column(parentId);"`
	MenuName		string	`orm:"column(menuName);"`
	MenuSort		int		`orm:"column(menuSort);"`
	DataFlag		int		`orm:"column(dataFlag);"`
	IsParent		bool	`orm:"-"`
}

func (this *Menus) TableName() string {
	return GetTableName("menus")
}

func (this *Menus) ListQuery() []*Menus {
	menus := make([]*Menus, 0)
	table := this.TableName()
	Db(table).Filter("dataFlag", 1).All(&menus)

	if len(menus) > 0 {
		for key, value := range menus {
			value.IsParent = true
			menus[key] = value
		}	

	}

	return menus

}


func (this *Menus) GetById(Id int) Menus{
	var menu Menus
	table := this.TableName()
	Db(table).Filter("menuId", Id).Filter("dataFlag", 1).One(&menu)
	return menu
}

