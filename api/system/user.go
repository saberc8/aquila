package system

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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
	// q: 为什么传的小写的username，然后req接收不到
	// a: 因为结构体中的字段名是大写的，所以无法接收到小写的username
	// q: 如何解决这个问题
	// a: 将结构体中的字段名改为小写
	fmt.Println(req.Username)
	// q: .Error 的返回 什么情况说明有用户，什么情况没有用户
	// a: 有用户的情况下，err为nil，没有用户的情况下，err不为nil
	var user model.UserEntity
	err = global.AquilaDb.Where("username = ?", req.Username).First(&user).Error

	if err != nil {
		// 创建新用户
		user = model.UserEntity{
			Username: req.Username,
			Password: req.Password,
		}
		err = global.AquilaDb.Create(&user).Error
		if err != nil {
			utils.Fail(ctx, "User creation failed")
			return
		}
		utils.Success(ctx, nil)
		return
	}
	utils.Fail(ctx, "User already exists")
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

func (u User) GetUserList(ctx *gin.Context) {

}
