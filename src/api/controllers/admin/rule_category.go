package admin

import (
	"api/controllers"
	"api/models"
	"strconv"
)

type RuleCategory struct {
	controllers.BaseController
}

var (
	ruleCategoryModel models.RuleCategoryModel
)

func (r_c RuleCategory) AddRuleCategory() {

	cate_name := r_c.GetString("cate_name")
	if cate_name == "" {
		r_c.JsonResult(1, "参数缺失")
		return
	}

	ruleCategoryModel.CategoryName = cate_name
	res := ruleCategoryModel.Insert(ruleCategoryModel)

	if res {
		r_c.JsonResult(0, "成功")
		return
	}
	r_c.JsonResult(1, "失败")
	return
}

// 获取权限分类
func (r_c RuleCategory) GetRuleCategory() {

	data := ruleCategoryModel.Select(0, 0)

	r_c.JsonResult(0, "成功", data)
	return
}

// 修改权限分类
func (r_c RuleCategory) EditRuleCategory() {
	id := r_c.GetString("id")
	val := r_c.GetString("val")

	if id == "" || val == "" {
		r_c.JsonResult(1, "参数缺失")
		return
	}

	intID, _ := strconv.Atoi(id)

	res := ruleCategoryModel.Edit(intID, val)
	if res {
		r_c.JsonResult(0, "成功")
		return
	}
	r_c.JsonResult(1, "失败")
	return
}

// 删除权限分类
func (r_c RuleCategory) DeleteRuleCategory() {
	id := r_c.GetString("id")

	intID, _ := strconv.Atoi(id)
	res := ruleCategoryModel.Delete(intID)

	if res {
		r_c.JsonResult(0, "成功")
		return
	}
	r_c.JsonResult(1, "失败")
	return
}
