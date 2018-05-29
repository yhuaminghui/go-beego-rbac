package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type AdminModel struct {
	Id         int
	Name       string
	Mobile     string
	Email      string
	Password   string
	CreateTime int
	UpdateTime int
}

func (this *AdminModel) Insert(param map[string]interface{}) int {
	t := time.Now()
	unix := t.Unix()
	sql := fmt.Sprintf("insert into admin(name,mobile,password,email,create_time,update_time) value(%q,%q,%q,%q,%d,%d)", param["name"], param["mobile"], param["password"], param["email"], unix, unix)

	fmt.Println(sql)
	fmt.Println(param)

	o := orm.NewOrm()
	num, err := o.Raw(sql).Exec()

	if err != nil {
		return 0
	}
	lastID, _ := num.LastInsertId()

	return int(lastID)
}

func (this *AdminModel) Select() []orm.Params {
	sql := "select admin.id,admin.name,admin.mobile,admin.email,admin.create_time,admin.is_on,group_concat(role.name) as role_name from admin left join admin_role on admin_role.admin_id = admin.id left join role on role.id = admin_role.role_id group by admin.id"

	o := orm.NewOrm()
	data := []orm.Params{}
	o.Raw(sql).Values(&data)

	return data
}

func (this *AdminModel) Find(id int) map[string]interface{} {
	sql := fmt.Sprintf("select * from admin where id = %d", id)

	o := orm.NewOrm()
	data := []orm.Params{}
	o.Raw(sql).Values(&data)

	return data[0]
}

// 修改管理员
func (this *AdminModel) UpdateAdmin(param map[string]interface{}, id int) bool {
	sql := fmt.Sprintf("update admin set name=%q,mobile=%q,email=%q,password=%q where id=%d", param["name"], param["mobile"], param["email"], param["password"], id)

	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}

func (this *AdminModel) UpdateIsOn(is_on int, id int) bool {
	sql := fmt.Sprintf("update admin set is_on = %d where id = %d", is_on, id)

	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}

func (this *AdminModel) Delete(id int) bool {
	sql := fmt.Sprintf("delete from admin where id = %d", id)

	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}
