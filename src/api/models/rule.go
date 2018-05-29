package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type RuleModel struct {
	Id             int `orm:auto`
	RuleCategoryID int
	Name           string
	Addr           string
	CreateTime     int
	UpdateTime     int
}

func (r RuleModel) Select(offset int, limit int) []orm.Params {

	where := ""
	if limit != 0 {
		where = fmt.Sprintf(" limit %d,%d", offset, limit)
	}

	sql := "select * from rule" + where

	var maps []orm.Params
	o := orm.NewOrm()
	o.Raw(sql).Values(&maps)

	return maps
}

func (r RuleModel) Insert(param RuleModel) bool {

	// 设置时间戳
	t := time.Now()
	intUnix := t.Unix()

	sql := fmt.Sprintf("insert into rule(rule_category_id,name,addr,create_time,update_time) value(%d,%q,%q,%d,%d)", param.RuleCategoryID, param.Name, param.Addr, intUnix, intUnix)

	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}

func (r RuleModel) Update(param map[string]string, id int) bool {
	t := time.Now()
	unix := t.Unix()

	sql := fmt.Sprintf("update rule set rule_category_id = %s,name = %q,addr = %q,update_time = %d where id = %d", param["category_id"], param["name"], param["addr"], unix, id)

	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()
	fmt.Println(err)
	fmt.Println(sql)
	if err != nil {
		return false
	}
	return true
}

func (r RuleModel) Delete(id int) bool {
	sql := fmt.Sprintf("delete from rule where id = %d", id)

	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}
