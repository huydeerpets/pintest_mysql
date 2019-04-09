package models

import (

	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id   int    
	RoleName string `orm:"unique;size(60)"`
	
}
const (
	USER_ROLE_READER = 0
	USER_ROLE_PUBLISHER = 1
	USER_ROLE_ADMIN = 2

)
// Insert role into database
func (r *Role) Insert() error {
	_, err := orm.NewOrm().Insert(r)
	return err
}

// Read role from database by given field
func (r *Role) Read(fields ...string) error {
	return orm.NewOrm().Read(r, fields...)
}

// Update role data in database
// if fields are provided, only this fields are updated
func (r *Role) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(r, fields...)
	return err
}

// Delete role from database
func (r *Role) Delete() error {
	_, err := orm.NewOrm().Delete(r)
	return err
}

// Roles is a helper to query the role table
func Roles() orm.QuerySeter {
	var table Role
	return orm.NewOrm().QueryTable(table)
}


