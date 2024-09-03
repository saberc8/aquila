package system

import (
	"aquila/global"
	"aquila/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u User) CreateUserApi(ctx *gin.Context) {
	var req UserDto
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("step1", err)
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}

	err = global.AquilaDb.Where("username = ?", req.Username).First(&req).Error

	if err != nil {
		utils.Fail(ctx, "User creation failed")
		return
	}
	utils.Success(ctx, "User created successfully")
}

func (u User) GetUserList(ctx *gin.Context) {
	
}
