package common

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-contrib/sessions"
)

type Auth struct{}

// Captcha 获取验证码
func (a Auth) Captcha(ctx *gin.Context) {
	svg, code := utils.GenerateSVG(80, 40)
	session := sessions.Default(ctx)
	session.Set("captcha", code)
	fmt.Println("captcha:", code)
	err := session.Save()
	if err != nil {
		fmt.Println("Session save error:", err)
		return
	}
	// 设置 Content-Type 为 "image/svg+xml"
	ctx.Header("Content-Type", "image/svg+xml; charset=utf-8")
	ctx.Data(http.StatusOK, "image/svg+xml", svg)
}

// Login 用户登录
func (a Auth) Login(ctx *gin.Context) {
	var req LoginDto
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "登录失败")
		return
	}
	fmt.Println(req)

	//session := sessions.Default(ctx)
	//fmt.Println("session:", session)
	//captcha := session.Get("captcha")
	//fmt.Println("captcha:", captcha)
	//if captcha == nil || captcha != req.Code {
	//	fmt.Println("captcha:", captcha)
	//	utils.Fail(ctx, "验证码错误")
	//	return
	//}
	var user model.UserEntity
	err = global.AquilaDb.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在")
		return
	}
	if user.Password != fmt.Sprintf("%x", md5.Sum([]byte(req.Password))) {
		utils.Fail(ctx, "密码错误")
		return
	}
	var LoginInfo LoginVo
	fmt.Println("user:", user.ID)
	LoginInfo.Token = utils.GenerateToken(int(user.ID))
	LoginInfo.UserInfo = UserVo{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
	}
	utils.Success(ctx, LoginInfo)
}

func (a Auth) Register(ctx *gin.Context) {
	var req RegisterDto
	err := ctx.ShouldBind(&req)
	fmt.Println(err)
	if err != nil {
		fmt.Println("step1", err)
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	session := sessions.Default(ctx)
	captcha := session.Get("captcha")
	if captcha == nil || captcha != req.Code {
		utils.Fail(ctx, "验证码错误")
		return
	}
	var user model.UserEntity
	// q: Find() 和 First() 的区别
	// a: Find() 返回的是一个数组，First() 返回的是一个对象
	err = global.AquilaDb.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		// 创建新用户
		user = model.UserEntity{
			Username: req.Username,
			Password: fmt.Sprintf("%x", md5.Sum([]byte(req.Password))),
		}
		err = global.AquilaDb.Transaction(func(tx *gorm.DB) error {
			err = tx.Create(&user).Error
			if err != nil {
				return err
			}
			utils.Success(ctx, nil)
			return nil
		})
		if err != nil {
			utils.Fail(ctx, "用户创建失败")
			return
		}
		utils.Success(ctx, nil)
		return
	} else {
		utils.Fail(ctx, "用户已存在")
		return
	}
}

func (a Auth) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	err := session.Save()
	if err != nil {
		utils.Fail(ctx, "退出登录失败")
		return
	}
	utils.Success(ctx, nil)
}
