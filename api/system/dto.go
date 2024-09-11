package system

import "aquila/model"

type UserDto struct {
	Id       int    `form:"id" json:"id"`
	Username string `form:"username" json:"username"` // 账号
	Password string `form:"password" json:"password"` // 密码
}

type RoleDto struct {
	Id     int    `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`     // 角色名称
	Remark string `form:"remark" json:"remark"` // 备注
	Status int64  `form:"status" json:"status"` // 状态
}

type RolePageDto struct {
	Total int64 `form:"total" json:"total"`
	Data  []model.RoleEntity
}

type MenuDto struct {
	Id        int    `form:"id" json:"id"`
	Name      string `form:"name" json:"name"`           // 菜单名称
	ParentId  int64  `form:"parentId" json:"parentId"`   // 父菜单ID
	OrderNum  int64  `form:"orderNum" json:"orderNum"`   // 排序
	Path      string `form:"path" json:"path"`           // 路由地址
	Component string `form:"component" json:"component"` // 组件路径
	Query     string `form:"query" json:"query"`         // 请求参数
	IsFrame   int64  `form:"isFrame" json:"isFrame"`     // 是否外链
	MenuType  string `form:"menuType" json:"menuType"`   // 菜单类型
	IsCatch   int64  `form:"isCatch" json:"isCatch"`     // 缓存
	IsHidden  int64  `form:"isHidden" json:"isHidden"`   // 是否隐藏
	Perms     string `form:"perms" json:"perms"`         // 权限标识
	Icon      string `form:"icon" json:"icon"`           // 图标
	Status    int64  `form:"status" json:"status"`       // 状态
	Remark    string `form:"remark" json:"remark"`       // 备注
}

type MenuPageDto struct {
	Total int64 `form:"total" json:"total"`
	Data  []model.MenuEntity
}

type RoleMenuDto struct {
	RoleId  int64  `form:"roleId" json:"roleId"`
	MenuIds string `form:"menuIds" json:"menuIds"`
}
