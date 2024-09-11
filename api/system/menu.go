package system

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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

func (m *Menu) GetMenuPageApi(ctx *gin.Context) {
	var req MenuPageDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	pageNum := ctx.DefaultQuery("pageNum", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	var menus []model.MenuEntity
	var total int64
	err := global.AquilaDb.Model(&model.MenuEntity{}).Count(&total).Error
	if err != nil {
		utils.Fail(ctx, "查询失败")
		return

	}
	err = global.AquilaDb.Scopes(utils.Paginate(pageNumInt, pageSizeInt)).Find(&menus).Error
	if err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}
	req.Total = total
	req.Data = menus
	utils.Success(ctx, req)
}
