package admin

import (
	"api/models"
	"strconv"

	"api/controllers"
)

type UsersController struct {
	controllers.BaseController
}

func (this *UsersController) AddUser() {
	email := this.GetString("email")
	password := this.GetString("pass")
	name := this.GetString("name")
	mobile := this.GetString("mobile")
	repass := this.GetString("repass")
	id := this.GetString("id")

	if password != repass {
		this.JsonResult(1, "两次密码不一致")
		return
	}

	keys := make([]string, 0)
	this.Ctx.Input.Bind(&keys, "role")

	var adminModel models.AdminModel
	maps := make(map[string]interface{}, 10)
	maps["name"] = name
	maps["email"] = email
	maps["mobile"] = mobile
	maps["password"] = password
	var insertID int
	if id != "" {
		intID, _ := strconv.Atoi(id)
		adminModel.UpdateAdmin(maps, intID)
		insertID = intID
	} else {
		insertID = adminModel.Insert(maps)
	}

	var res bool
	if id != "" {
		// 删除之前用户的角色
		var adminRoleModel models.AdminRoleModel
		adminRoleModel.Delete(insertID)
	}

	adminRoleData := make([]map[string]int, len(keys))
	for k, v := range keys {

		tempMaps := make(map[string]int)

		intRoleID, _ := strconv.Atoi(v)

		tempMaps["role_id"] = intRoleID
		tempMaps["admin_id"] = insertID

		adminRoleData[k] = tempMaps
	}

	var adminRoleModel models.AdminRoleModel
	res = adminRoleModel.Insert(adminRoleData)

	if res {
		this.JsonResult(0, "成功")
		return
	}
	this.JsonResult(1, "失败")
	return
}

// 获取管理员列表
func (this *UsersController) GetAdminList() {
	var userModel models.AdminModel
	data := userModel.Select()

	this.JsonResult(0, "成功", data)
	return
}

// 启用管理员
func (this *UsersController) StartAdmin() {
	isOn := this.GetString("is_on")
	id := this.GetString("id")

	// 处理参数
	intIsOn, _ := strconv.Atoi(isOn)
	intID, _ := strconv.Atoi(id)

	var adminModel models.AdminModel
	res := adminModel.UpdateIsOn(intIsOn, intID)

	if res {
		this.JsonResult(0, "成功")
		return
	}
	this.JsonResult(1, "失败")
	return
}

// 删除管理员
func (this *UsersController) DeleteAdmin() {
	id := this.GetString("id")

	// 处理参数
	intID, _ := strconv.Atoi(id)

	var adminModel models.AdminModel
	res := adminModel.Delete(intID)

	if res {
		this.JsonResult(0, "成功")
		return
	}
	this.JsonResult(1, "失败")
	return
}

// 编辑管理员 回显
func (this *UsersController) GetAdmin() {
	id := this.GetString("id")
	intID, _ := strconv.Atoi(id)

	// 获取管理员
	var adminModel models.AdminModel
	adminData := adminModel.Find(intID)

	// 获取当前管理员的角色
	var adminRoleModel models.AdminRoleModel
	adminRoleData := adminRoleModel.Find(intID)

	adminData["role"] = adminRoleData

	this.JsonResult(0, "成功", adminData)
	return
}
