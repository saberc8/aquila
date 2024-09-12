package system

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type User struct {
}

func (u User) CreateUserApi(ctx *gin.Context) {
	var req UserDto
	err := ctx.ShouldBind(&req)
	fmt.Println(err)
	if err != nil {
		fmt.Println("step1", err)
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	fmt.Printf("Received request: %+v\n", req) // 打印接收到的请求内容
	var user model.UserEntity
	err = global.AquilaDb.Where("username = ?", req.Username).First(&user).Error

	if err != nil {
		// 创建新用户
		user = model.UserEntity{
			Username: req.Username,
			Nickname: req.Nickname,
			Password: fmt.Sprintf("%x", md5.Sum([]byte(req.Password))),
		}
		err = global.AquilaDb.Create(&user).Error
		if err != nil {
			utils.Fail(ctx, "用户创建失败")
			return
		}
		utils.Success(ctx, nil)
		return
	}
	utils.Fail(ctx, "用户已存在")
}

func (u User) GetUserApi(ctx *gin.Context) {
	var req UserDto
	// get请求参数绑定
	err := ctx.ShouldBind(&req)
	fmt.Println(err)
	if err != nil {
		fmt.Println("step1", err)
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	var user model.UserEntity
	fmt.Printf("Received request: %+v\n", req) // 打印接收到的请求内容
	if req.Username == "" {

		var uid, _ = ctx.Get("uid")
		err = global.AquilaDb.Where("id = ?", uid).Find(&user).Error
		if err != nil {
			utils.Fail(ctx, "用户不存在")
			return
		}
		utils.Success(ctx, user)
		return
	}
	err = global.AquilaDb.Where("username = ?", req.Username).Find(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在")
		return
	}
	utils.Success(ctx, user)
}

func (u User) ChangePasswordApi(ctx *gin.Context) {
	var req ChangePasswordDto
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}

	var user model.UserEntity
	err = global.AquilaDb.Where("id = ?", req.UserId).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在")
		return
	}
	if user.Password != fmt.Sprintf("%x", md5.Sum([]byte(req.OldPassword))) {
		utils.Fail(ctx, "旧密码错误")
		return
	}

	err = global.AquilaDb.Model(&user).Update("password", fmt.Sprintf("%x", md5.Sum([]byte(req.NewPassword)))).Error
	if err != nil {
		utils.Fail(ctx, "修改失败")
		return
	}
	utils.Success(ctx, nil)
}
func (u User) GetUserPageApi(ctx *gin.Context) {
	var req UserDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}

	pageNum := ctx.DefaultQuery("pageNum", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	var users []model.UserEntity
	var total int64
	var res UserPageVo
	// q: 如何按照id顺序下发数据
	// a: global.AquilaDb.Order("id desc").Scopes(utils.Paginate(pageNumInt, pageSizeInt)).Find(&users)
	global.AquilaDb.Model(&model.UserEntity{}).Count(&total)
	global.AquilaDb.Order("id asc").Scopes(utils.Paginate(pageNumInt, pageSizeInt)).Find(&users)

	res.Total = total
	res.Data = users

	utils.Success(ctx, res)
}

func (u User) UpdateUserApi(ctx *gin.Context) {
	var req UserDto
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}
	var user model.UserEntity
	err = global.AquilaDb.Where("id = ?", req.Id).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在")
		return
	}
	user.Username = req.Username
	user.Nickname = req.Nickname
	err = global.AquilaDb.Save(&user).Error
	if err != nil {
		utils.Fail(ctx, "更新失败")
		return
	}
	utils.Success(ctx, nil)
}

func (u User) BindRoleApi(ctx *gin.Context) {
	var req UserRoleDto
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}
	fmt.Println(req, req.UserId)
	var userRoles []model.UserRoleEntity
	var userRole model.UserRoleEntity
	roleIds := req.RoleIds
	roleIdArr := utils.StrSplit(roleIds)
	err = global.AquilaDb.Transaction(func(tx *gorm.DB) error {
		err = tx.Unscoped().Where("user_id = ?", req.UserId).Delete(&userRole).Error
		if err != nil {
			return err
		}
		for _, v := range roleIdArr {
			roleId, _ := strconv.Atoi(strconv.Itoa(v))
			// 查询角色是否存在
			var role model.RoleEntity
			err = tx.Where("id = ?", roleId).First(&role).Error
			if err == nil {
				userRole = model.UserRoleEntity{
					UserId: req.UserId,
					RoleId: int64(roleId),
				}
				userRoles = append(userRoles, userRole)
			}
		}
		err = tx.Create(&userRoles).Error
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

func (u User) GetUserMenuApi(ctx *gin.Context) {
	var req GetUserMenuDto
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}
	var user model.UserEntity

	if req.UserId == 0 {
		var uid, _ = ctx.Get("uid")
		err = global.AquilaDb.Where("id = ?", uid).Find(&user).Error
		if err != nil {
			utils.Fail(ctx, "用户不存在")
			return
		}
	} else {
		err = global.AquilaDb.Where("id = ?", req.UserId).Find(&user).Error
		if err != nil {
			utils.Fail(ctx, "用户不存在")
			return
		}
	}
	var userRoles []model.UserRoleEntity
	err = global.AquilaDb.Where("user_id = ?", user.ID).Find(&userRoles).Error
	if err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}
	var roleIds []int
	for _, userRole := range userRoles {
		roleIds = append(roleIds, int(userRole.RoleId))
	}
	var roleMenus []model.RoleMenuEntity
	err = global.AquilaDb.Where("role_id in (?)", roleIds).Find(&roleMenus).Error
	if err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}
	var menuIds []int
	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, int(roleMenu.MenuId))
	}
	var menus []model.MenuEntity
	err = global.AquilaDb.Where("id in (?)", menuIds).Find(&menus).Error
	if err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}
	// = 0 返回空数组
	if len(menus) == 0 {
		utils.Success(ctx, menus)
		return
	}
	// menus根据id和parentId，组装成树形结构，children: []MenuEntity
	var menuTree []UserMenuTreeDto
	menuTree = getMenuTree(0, menus)
	utils.Success(ctx, menuTree)
}

func getMenuTree(parentId int64, menus []model.MenuEntity) []UserMenuTreeDto {
	var menuTree []UserMenuTreeDto
	for _, menu := range menus {
		if menu.ParentId == parentId {
			children := getMenuTree(menu.ID, menus)
			menuTree = append(menuTree, UserMenuTreeDto{
				MenuDto: MenuDto{
					Id:        int(menu.ID),
					Name:      menu.Name,
					ParentId:  menu.ParentId,
					OrderNum:  menu.OrderNum,
					Path:      menu.Path,
					Component: menu.Component,
					Query:     menu.Query,
					IsFrame:   menu.IsFrame,
					MenuType:  menu.MenuType,
					IsCatch:   menu.IsCatch,
					IsHidden:  menu.IsHidden,
					Perms:     menu.Perms,
					Icon:      menu.Icon,
					Status:    menu.Status,
					Remark:    menu.Remark,
				},
				Children: children,
			})
		}
	}
	return menuTree
}
