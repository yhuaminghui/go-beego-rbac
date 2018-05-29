// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {

	// 添加管理员
	beego.Router("/add_user", &admin.UsersController{}, "post:AddUser")
	// 获取管理员列表
	beego.Router("/get_admin_list", &admin.UsersController{}, "get:GetAdminList")
	// 启用管理员
	beego.Router("/start_admin", &admin.UsersController{}, "post:StartAdmin")
	// 删除管理员
	beego.Router("/delete_admin", &admin.UsersController{}, "post:DeleteAdmin")
	// 编辑 管理员 回显
	beego.Router("/get_admin", &admin.UsersController{}, "get:GetAdmin")
	// 添加权限分类
	beego.Router("/add_rule_category", &admin.RuleCategory{}, "post:AddRuleCategory")
	// 获取权限列表分类
	beego.Router("/get_rule_category", &admin.RuleCategory{}, "get:GetRuleCategory")
	// 修改权限分类
	beego.Router("/edit_rule_category", &admin.RuleCategory{}, "post:EditRuleCategory")
	// 删除权限分类
	beego.Router("/delete_rule_category", &admin.RuleCategory{}, "post:DeleteRuleCategory")

	// 规则列表数据
	beego.Router("/get_rule_list_and_rule_category", &admin.RuleController{}, "get:GetRuleListData")
	// 添加规则
	beego.Router("/add_rule", &admin.RuleController{}, "post:AddRule")
	// 修改规则
	beego.Router("/edit_rule", &admin.RuleController{}, "post:EditRule")
	// 删除规则
	beego.Router("/delete_rule", &admin.RuleController{}, "post:DeleteRUle")

	// 获取添加角色时的权限数据
	beego.Router("/get_rule_form_rule_add", &admin.RoleAddController{}, "get:GetRule")
	// 添加角色
	beego.Router("/add_role", &admin.RoleAddController{}, "post:AddRole")
	// 获取角色
	beego.Router("get_role", &admin.RoleController{}, "get:GetRole")
	// 启用角色
	beego.Router("/start_role", &admin.RoleController{}, "post:StartRole")
	// 删除角色
	beego.Router("/delete_role", &admin.RoleController{}, "post:DeleteRole")
	// 修改角色 回显数据
	beego.Router("/get_role_and_rule", &admin.RoleController{}, "get:GetRoleAndRule")

	// test
	beego.Router("/test", &admin.IndexController{}, "get:Say")
}
