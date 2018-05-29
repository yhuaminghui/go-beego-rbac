package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type AdminRoleModel struct {
	Id         int
	RoleID     int
	AdminID    int
	CreateTime int
	UpdateTime int
}

func (this *AdminRoleModel) Insert(params []map[string]int) bool {
	t := time.Now()
	unix := t.Unix()

	o := orm.NewOrm()
	var errs error
	for _, v := range params {
		sql := fmt.Sprintf("insert into admin_role(role_id,admin_id,create_time,update_time) value(%d,%d,%d,%d)", v["role_id"], v["admin_id"], unix, unix)
		_, err := o.Raw(sql).Exec()
		errs = err
	}

	if errs != nil {
		return false
	}
	return true
}

func (this *AdminRoleModel) Find(id int) []orm.Params {
	sql := fmt.Sprintf("select * from admin_role where admin_id = %d", id)

	o := orm.NewOrm()
	data := []orm.Params{}
	o.Raw(sql).Values(&data)

	return data
}

func (this *AdminRoleModel) Delete(id int) bool {
	sql := fmt.Sprintf("delete from admin_role where id = %d", id)

	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}
