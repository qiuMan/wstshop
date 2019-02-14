package models

// import "reflect"

 type Privileges struct {
	 Id					int 	  `orm:"pk;column(privilegeId);"`
	 MenuId				int		  `orm:"column(menuId);"`
	 PrivilegeCode		string	  `orm:"column(privilegeCode);"`
	 PrivilegeName		string	  `orm:"column(privilegeName);"`
	 IsMenuPrivilege	int8 	  `orm:"column(isMenuPrivilege);"`
	 PrivilegeUrl		string 	  `orm:"column(privilegeUrl);"`
     OtherPrivilegeUrl	string    `orm:"column(otherPrivilegeUrl);"`
	 DataFlag			int8    `orm:"column(dataFlag);"`
 }

func NewPrivileges() *Privileges {
	obj := new(Privileges)
	return obj
}

 func (this *Privileges) TableName() string {
    // db table name
    return GetTableName("privileges")
}

 func (this *Privileges) ListQuery(id int) []*Privileges  {
	table := this.TableName()
	data := make([]*Privileges, 0)
	Db(table).Filter("menuId", id).Filter("DataFlag", 1).All(&data)
	return data

	//  fmt.Println(privilege)
 }

 func (this *Privileges) GetById(Id int) Privileges {
	var privileges Privileges
	table := this.TableName()
	Db(table).Filter("privilegeId", Id).Filter("dataFlag", 1).One(&privileges)
	return privileges
 }

 func (this *Privileges) Add(data *Privileges) (int64, error) {
	table := this.TableName()
	db := Db(table)
	i, _ := db.PrepareInsert()
	id, err := i.Insert(data)	 
	
	return id, err
 }

func (this *Privileges) Update(filter map[string]interface{}, data map[string]interface{}) (int64, error) {
	table := this.TableName()
	db := Db(table)
	for key, value := range filter {
		db = db.Filter(key, value)
	}
	
	num, err := db.Update(data)
	return num, err
}

func (this *Privileges) Delete(id int) (int64, error){
	table := this.TableName()
	db := Db(table)
	num, err := db.Filter("privilegeId", id).Delete()
	return num, err
}