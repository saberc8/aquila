package system

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Role struct{}

func (r Role) CreateRoleApi(ctx *gin.Context) {
	var req RoleDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	var role model.RoleEntity
	err := global.AquilaDb.Where("name = ?", req.Name).First(&role).Error
	if err != nil {
		role = model.RoleEntity{
			Name:   req.Name,
			Remark: req.Remark,
			Status: req.Status,
		}
		err = global.AquilaDb.Create(&role).Error
		if err != nil {
			utils.Fail(ctx, "Role creation failed")
			return
		}
		utils.Success(ctx, nil)
		return
	}
	utils.Fail(ctx, "角色已存在")
}

func (r Role) GetRoleApi(ctx *gin.Context) {
	var req RoleDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
}

// updateRoleApi 更新角色
func (r Role) UpdateRoleApi(ctx *gin.Context) {
	var req RoleDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	var role model.RoleEntity
	err := global.AquilaDb.Where("id = ?", req.Id).First(&role).Error
	if err != nil {
		utils.Fail(ctx, "角色不存在")
		return
	}
	role.Name = req.Name
	role.Remark = req.Remark
	role.Status = req.Status
	err = global.AquilaDb.Save(&role).Error
	if err != nil {
		utils.Fail(ctx, "角色更新失败")
		return
	}
	utils.Success(ctx, nil)
}

func (r Role) GetRolePageApi(ctx *gin.Context) {
	var req RolePageDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}

	pageNum := ctx.DefaultQuery("pageNum", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	var roles []model.RoleEntity
	var total int64

	global.AquilaDb.Model(&model.RoleEntity{}).Count(&total)
	global.AquilaDb.Scopes(utils.Paginate(pageNumInt, pageSizeInt)).Find(&roles)

	req.Total = total
	req.Data = roles

	utils.Success(ctx, req)
}

// BindMenuApi 绑定菜单
func (r Role) BindMenuApi(ctx *gin.Context) {
	var req RoleMenuDto
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}
	fmt.Println(req, req.RoleId)
	var role model.RoleEntity
	err = global.AquilaDb.Where("id = ?", req.RoleId).First(&role).Error
	if err != nil {
		utils.Fail(ctx, "角色不存在")
		return
	}
	var roleMenus []model.RoleMenuEntity
	var roleMenu model.RoleMenuEntity
	menus := req.MenuIds
	fmt.Println(menus)
	err = global.AquilaDb.Transaction(func(tx *gorm.DB) error {
		err = tx.Unscoped().Where("role_id = ?", req.RoleId).Delete(&roleMenu).Error
		if err != nil {
			return err
		}
		for _, v := range utils.StrSplit(menus) {
			menuId, _ := strconv.Atoi(strconv.Itoa(v))
			var menu model.MenuEntity
			err = tx.Where("id = ?", menuId).First(&menu).Error
			if err == nil {
				roleMenu = model.RoleMenuEntity{
					RoleId: req.RoleId,
					MenuId: int64(menuId),
				}
				roleMenus = append(roleMenus, roleMenu)
			}
		}
		err = tx.Create(&roleMenus).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		utils.Fail(ctx, "绑定失败"+err.Error())
		return
	}
	utils.Success(ctx, nil)
}
