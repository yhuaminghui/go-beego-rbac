package admin

import (
	"api/controllers"
	"api/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type RuleController struct {
	controllers.BaseController
}

var (
	ruleModel models.RuleModel
)

func (r RuleController) GetRuleListData() {
	// 获取权限分类
	var ruleCategoryModel models.RuleCategoryModel
	ruleCategoryInfo := ruleCategoryModel.Select(0, 0)

	// 获取权限列表数据
	var ruleModel models.RuleModel
	ruleInfo := ruleModel.Select(0, 0)

	for k, v := range ruleInfo {
		for _, value := range ruleCategoryInfo {
			if v["rule_category_id"] == value["id"] {
				ruleInfo[k]["category_name"] = value["category_name"]
			}
		}
	}
	returnData := make(map[string][]orm.Params)
	returnData["rule_category"] = ruleCategoryInfo
	returnData["rule"] = ruleInfo

	// fmt.Println(returnData)

	r.JsonResult(0, "成功", returnData)
	return
}

func (this RuleController) AddRule() {
	categoryID := this.GetString("category_id")
	name := this.GetString("name")
	addr := this.GetString("addr")
	fmt.Println("中国")
	// 验证
	if categoryID == "" || name == "" || addr == "" {
		this.JsonResult(1, "参数缺失")
		return
	}
	intCategoryID, err := strconv.Atoi(categoryID)
	if err != nil {
		this.JsonResult(1, "参数：规则类型不是数字")
		return
	}

	// 插入数据
	ruleModel.Name = name
	ruleModel.Addr = addr

	ruleModel.RuleCategoryID = intCategoryID
	isSuccess := ruleModel.Insert(ruleModel)

	if isSuccess {
		this.JsonResult(0, "成功")
		return
	}
	this.JsonResult(1, "失败")
}

// 修改
func (r RuleController) EditRule() {
	categoryID := r.GetString("category_id")
	addr := r.GetString("addr")
	name := r.GetString("name")
	intID, _ := strconv.Atoi(r.GetString("id"))

	if intID == 0 {
		r.JsonResult(1, "ID缺失")
		return
	}

	params := make(map[string]string)
	params["category_id"] = categoryID
	params["addr"] = addr
	params["name"] = name

	isSuccess := ruleModel.Update(params, intID)
	if isSuccess {
		r.JsonResult(0, "成功")
		return
	}
	r.JsonResult(1, "失败")
	return
}

// 删除
func (r *RuleController) DeleteRUle() {
	intID, err := strconv.Atoi(r.GetString("id"))

	if err != nil {
		r.JsonResult(1, "ID有误")
		return
	}

	isSuccess := ruleModel.Delete(intID)

	if isSuccess {
		r.JsonResult(0, "成功")
		return
	}
	r.JsonResult(1, "失败")
	return
}
