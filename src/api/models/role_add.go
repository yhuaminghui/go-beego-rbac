package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type RoleModel struct {
	Id         int
	Name       string
	Discribe   string
	CreateTime int
	UpdateTime int
}

func (r RoleModel) Select() []orm.Params {
	sql := "select role.id,role.name,role.discribe,role.is_on,GROUP_CONCAT(rule.name) as rule_name from role left join role_rule on role.id = role_rule.role_id left join rule on role_rule.rule_id = rule.id GROUP BY role.id"
	o := orm.NewOrm()
	data := []orm.Params{}
	o.Raw(sql).Values(&data)

	return data
}

// 插入数据
func (r *RoleModel) Insert(param map[string]interface{}) {

	t := time.Now()
	unix := t.Unix()

	sql := fmt.Sprintf("insert into role(name,discribe,create_time,update_time) value(%q,%q,%d,%d)", param["name"], param["desc"], unix, unix)
	fmt.Println(sql)
	o := orm.NewOrm()
	num, _ := o.Raw(sql).Exec()

	lastID, _ := num.LastInsertId()

	ruleID := param["rule_id"]
	slice := ruleID.([]string)

	for _, v := range slice {
		sql1 := fmt.Sprintf("insert into role_rule(role_id,rule_id,create_time,update_time) values(%d,%s,%d,%d)", lastID, v, unix, unix)
		fmt.Println(sql1)
		o.Raw(sql1).Exec()
	}

}

func (this *RoleModel) UpdateIsOn(is_on int, id int) bool {
	sql := fmt.Sprintf("update role set is_on = %d where id = %d", is_on, id)

	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}

func (this *RoleModel) Delete(id int) bool {
	sql := fmt.Sprintf("delete from role where id = %d", id)

	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}

// 修改角色 回显数据
func (this *RoleModel) EditForGet(id int) ([]orm.Params, []orm.Params) {
	o := orm.NewOrm()

	roleSql := fmt.Sprintf("select * from role where id = %d", id)
	roleData := []orm.Params{}
	o.Raw(roleSql).Values(&roleData)

	ruleSql := fmt.Sprintf("select * from role_rule left join rule on rule.id = role_rule.rule_id where role_rule.role_id = %d", id)
	ruleData := []orm.Params{}
	o.Raw(ruleSql).Values(&ruleData)

	return roleData, ruleData
}

// 修改角色信息
func (this *RoleModel) UpdateRole(param map[string]interface{}, id int) bool {
	roleSql := fmt.Sprintf("update role set name = %q,discribe = %q where id = %d", param["name"], param["desc"], id)

	o := orm.NewOrm()
	_, err := o.Raw(roleSql).Exec()

	if err != nil {
		return false
	}

	// 修改角色权限
	// 先删除当前角色的所有权限
	roleRuleDelSql := fmt.Sprintf("delete from role_rule where role_id = %d", id)
	o.Raw(roleRuleDelSql).Exec()
	// 添加权限
	t := time.Now()
	unix := t.Unix()
	roleRUle := param["rule_id"].([]string)

	var errs error
	for _, v := range roleRUle {
		roleRuleAddSql := fmt.Sprintf("insert into role_rule(role_id,rule_id,create_time,update_time) values(%d,%s,%d,%d)", id, v, unix, unix)
		_, err := o.Raw(roleRuleAddSql).Exec()
		errs = err
	}

	if errs != nil {
		return false
	}
	return true
}
