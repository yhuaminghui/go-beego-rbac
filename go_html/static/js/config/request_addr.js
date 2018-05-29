// 基础服务器地址
var serverAddr = 'http://beego.golang.com:8080'; 

/*********************** 服务器请求 begin ****************************/
// 添加管理员
var serverRegisterAddr = serverAddr + '/add_user'
// 获取管理员列表
var serverGetAdminList = serverAddr + '/get_admin_list'
// 修改管理员为启用状态
var serverStartAdmin = serverAddr + '/start_admin'
// 删除管理员
var serverDeleteAdmin = serverAddr + '/delete_admin'
// 编辑管理员 回显数据
var serverGetAdmin = serverAddr + '/get_admin'

// 添加权限
var serverAddRule = serverAddr + '/add_rule' 
// 修改 权限
var serverEditRule = serverAddr + '/edit_rule'
// 删除 权限
var serverDeleteRule = serverAddr + '/delete_rule'

// 添加权限分类 
var ruleCategory = serverAddr + '/add_rule_category'
// 查询权限分类
var getRuleCategory = serverAddr + '/get_rule_category'
// 修改权限分类 
var editRuleCategory = serverAddr + '/edit_rule_category'
// 删除权限分类
var deleteRuleCategory = serverAddr + '/delete_rule_category'

// 获取权限列表数据与权限分类
var ruleListAndRuleCategory = serverAddr + '/get_rule_list_and_rule_category'

// 获取添加角色页面的权限数据
var serverGetRuleFormRoleAdd = serverAddr + '/get_rule_form_rule_add'
// 添加角色
var serverAddRole = serverAddr + '/add_role'
// 获取角色
var serverRoleList = serverAddr + '/get_role'
// 启用角色
var serverStartRole = serverAddr + '/start_role'
// 删除角色
var serverDeleteRole = serverAddr + '/delete_role'
// 修改角色 获取角色对应信息回显
var serverGetRoleAndRule = serverAddr + '/get_role_and_rule'


// 首页地址
var serverIndexAddr = serverAddr + '/index';

/*********************** 服务器请求 end ****************************/