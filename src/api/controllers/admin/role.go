package admin

import (
	"api/controllers"
	"api/models"
	"strconv"
)

type RoleController struct {
	controllers.BaseController
}

func (r *RoleController) GetRole() {
	// 获取当前的角色信息与角色对应的权限
	var roleModel models.RoleModel
	roleData := roleModel.Select()

	r.JsonResult(0, "成功", roleData)
	return

}

func (this *RoleController) StartRole() {
	id := this.GetString("id")
	intID, _ := strconv.Atoi(id)
	isOn := this.GetString("is_on")
	intIsOn, _ := strconv.Atoi(isOn)

	var roleModel models.RoleModel
	res := roleModel.UpdateIsOn(intIsOn, intID)

	if res {
		this.JsonResult(0, "成功")
		return
	}
	this.JsonResult(1, "失败")
	return
}

func (this *RoleController) DeleteRole() {
	id := this.GetString("id")
	intID, _ := strconv.Atoi(id)

	var roleModel models.RoleModel
	res := roleModel.Delete(intID)

	if res {
		this.JsonResult(0, "成功")
		return
	}
	this.JsonResult(1, "失败")
	return
}

// 修改角色 回显数据
func (this *RoleController) GetRoleAndRule() {
	id := this.GetString("id")
	intID, _ := strconv.Atoi(id)

	var roleModel models.RoleModel
	roleData, ruleData := roleModel.EditForGet(intID)

	returnData := roleData[0]
	returnData["rule"] = ruleData

	this.JsonResult(0, "成功", returnData)
	return
}
