package system

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Menu struct{}

func (m *Menu) CreateMenuApi(ctx *gin.Context) {
	var req MenuDto
	err := ctx.ShouldBind(&req)
	fmt.Println(err)
	if err != nil {
		fmt.Println("step1", err)
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	fmt.Printf("Received request: %+v\n", req) // 打印接收到的请求内容
	var menu model.MenuEntity
	err = global.AquilaDb.Where("name = ?", req.Name).First(&menu).Error
	if err != nil {
		// 创建新菜单
		menu = model.MenuEntity{
			Name:      req.Name,
			ParentId:  req.ParentId,
			OrderNum:  req.OrderNum,
			Path:      req.Path,
			Component: req.Component,
			Query:     req.Query,
			IsFrame:   req.IsFrame,
			MenuType:  req.MenuType,
			IsCatch:   req.IsCatch,
			IsHidden:  req.IsHidden,
			Perms:     req.Perms,
			Icon:      req.Icon,
			Status:    req.Status,
			Remark:    req.Remark,
		}
		err = global.AquilaDb.Create(&menu).Error
		if err != nil {
			utils.Fail(ctx, "菜单创建失败")
			return
		}
		utils.Success(ctx, "菜单创建成功")
		return
	}
	utils.Fail(ctx, "菜单已存在")
}

func (m *Menu) GetMenuApi(ctx *gin.Context) {
	var req MenuDto
	// get请求参数绑定
	err := ctx.ShouldBind(&req)
	fmt.Println(err)
	if err != nil {
		fmt.Println("step1", err)
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	var menu model.MenuEntity
	err = global.AquilaDb.Where("name = ?", req.Name).First(&menu).Error
	if err != nil {
		utils.Fail(ctx, "菜单不存在")
		return
	}
	utils.Success(ctx, menu)
}

func (m *Menu) GetMenuAllApi(ctx *gin.Context) {
	var req MenuPageDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	var menus []model.MenuEntity
	// 获取所有的menu
	err := global.AquilaDb.Find(&menus).Error
	if err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}
	// menus 根据id和parentId，组装成树形结构，children: []MenuEntity
	var menuTree []UserMenuTreeDto
	menuTree = getMenuTree(0, menus)

	req.Data = menuTree
	req.Total = int64(len(menuTree))
	utils.Success(ctx, req)
}
