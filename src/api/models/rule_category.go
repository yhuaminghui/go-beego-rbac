package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type RuleCategoryModel struct {
	Id           int    `orm:"auto"`
	CategoryName string `orm:"column(category_name)"`
	CreateTime   int    `orm:"column(create_time)"`
	UpdateTime   int    `orm:"column(update_time)"`
}

// 获取
/**
@param id 权限分类ID
*/
func (r_c RuleCategoryModel) Select(offset int, limit int) []orm.Params {
	where := ""
	if limit != 0 {
		where = fmt.Sprintf(" limit %d,%d", offset, limit)

	}
	sql := "select * from rule_category" + where

	// var ruleCategoryData []RuleCategoryModel

	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(sql).Values(&maps)

	return maps
}

func (r_c RuleCategoryModel) Insert(ruleCategory RuleCategoryModel) bool {
	o := orm.NewOrm()

	t := time.Now()

	sql := fmt.Sprintf("insert into rule_category(category_name,create_time,update_time) values(%q,%d,%d)", ruleCategory.CategoryName, t.Unix(), t.Unix())

	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}

func (r_c RuleCategoryModel) Edit(id int, category_name string) bool {
	sql := fmt.Sprintf("update rule_category set category_name = %q where id = %d", category_name, id)
	fmt.Println(sql)
	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()
	if err != nil {
		return false
	}
	return true
}

func (r_c RuleCategoryModel) Delete(id int) bool {
	sql := fmt.Sprintf("delete from rule_category where id = %d", id)

	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}

// 获取一个
func (r_c RuleCategoryModel) Get(id int) interface{} {
	sql := fmt.Sprintf("select * from rule_category where id = %d", id)

	var ruleCategoryData RuleCategoryModel

	o := orm.NewOrm()
	err := o.Raw(sql).QueryRow(&ruleCategoryData)

	if err != nil {
		return false
	}
	return ruleCategoryData
}
