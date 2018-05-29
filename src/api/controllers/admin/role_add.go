package admin

import (
	"api/controllers"
	"api/models"
	"fmt"
	"strconv"
)

type RoleAddController struct {
	controllers.BaseController
}

func (r RoleAddController) GetRule() {
	// 获取权限
	// 权限表 模型
	var ruleModel models.RuleModel
	// 权限分类 模型
	var ruleCategory models.RuleCategoryModel

	// 获取素有权限
	ruleInfo := ruleModel.Select(0, 0)
	// 获取权限分类
	ruleCategoryInfo := ruleCategory.Select(0, 0)

	// 计算需要的内存大小
	ruleCategoryMemSize := len(ruleCategoryInfo)
	ruleMemSize := len(ruleInfo)

	ruleCategoryMem := make([]map[string]interface{}, ruleCategoryMemSize)

	var flag int
	for k, v := range ruleCategoryInfo {
		flag = 0

		maps := make(map[string]interface{})
		children := make([]interface{}, ruleMemSize)

		for _, value := range ruleInfo {
			if v["id"] == value["rule_category_id"] {
				children[flag] = value
				flag++
			}
		}

		maps["children"] = &children
		maps["id"] = v["id"]
		maps["name"] = v["category_name"]
		ruleCategoryMem[k] = maps
	}

	r.JsonResult(0, "成功", ruleCategoryMem)
	return
}

// 添加角色
func (r *RoleAddController) AddRole() {
	name := r.GetString("name")
	desc := r.GetString("desc")
	id := r.GetString("id")
	fmt.Println(id == "")

	keys := make([]string, 0)
	r.Ctx.Input.Bind(&keys, "rule_id")

	maps := make(map[string]interface{})
	maps["rule_id"] = keys
	maps["name"] = name
	maps["desc"] = desc

	roleModel := new(models.RoleModel)

	if id == "" {
		roleModel.Insert(maps)
	} else {
		intID, _ := strconv.Atoi(id)
		roleModel.UpdateRole(maps, intID)
	}

	r.JsonResult(0, "成功")
	return
}
