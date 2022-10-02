package controller

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/service"
	"GoLog/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户相关接口,注册,登录,签发token等

type UserController struct{}

func (u *UserController) PostSignUp(c *gin.Context) {
	var user model.RegisterUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		commen.FailedWithMsg(err.Error(), c)
	}
	// 绑定成功后,调用注册服务
	ruser := model.User{
		UserName: user.UserName,
		Password: user.Password,
		NickName: user.NickName,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
	regUserInfo, err := service.AllServiceApp.UserService.RegisterUser(ruser)
	if err != nil {
		commen.GVA_LOG.Error("注册失败", zap.Error(err))
		commen.FailedWithDetailed(regUserInfo, "注册失败", c)
		return
	}
	commen.OkWithDetailed(regUserInfo, "注册成功", c)
}

func (u *UserController) PostLogIn(c *gin.Context) {
	var user model.LogInUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		commen.FailedWithMsg(err.Error(), c)
	}

	// 绑定成功后,调用登录服务
	ruser := model.User{
		UserName: user.UserName,
		Password: user.Password,
	}
	loginUser, err := service.AllServiceApp.UserService.LogInUser(&ruser)
	if err != nil {
		commen.GVA_LOG.Error("登录失败", zap.Error(err))
		commen.FailedWithDetailed(loginUser, "登录失败", c)
		return
	}
	// 登录成功之后,签发JWT
	u.GetToken(c, *loginUser)
}

func (u *UserController) GetToken(c *gin.Context, user model.User) {
	j := utils.JWT{SigningKey: []byte(commen.GVA_CONFIG.JWT.SigningKey)}
	baseClaims := model.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
	}
	claims := j.SetClaims(baseClaims)
	token, err := j.GenToken(claims)
	if err != nil {
		commen.GVA_LOG.Error("设置登录状态失败", zap.Error(err))
		commen.FailedWithMsg("设置登录状态失败", c)
		return
	}
	commen.OkWithDetailed(model.LogInUserResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "设置登录状态成功", c)
}

func (u *UserController) ChangePwd(c *gin.Context) {
	var cp model.ChangePassword
	if err := c.ShouldBindJSON(&cp); err != nil {
		commen.FailedWithMsg(err.Error(), c)
	}

	// 绑定数据后,调用修改密码服务,首先要找到用户的ID
	userId := utils.GetUserID(c)
	user := &model.User{GVA_MODEL: commen.GVA_MODEL{ID: userId}, Password: cp.Password}
	// service.AllServiceApp.UserService.ChangePwd()
	_, err := service.AllServiceApp.UserService.ChangePwd(user, cp.NewPassword)
	if err != nil {
		commen.FailedWithMsg("修改密码失败,原密码与当前账户不符", c)
		commen.GVA_LOG.Error("修改密码失败!", zap.Error(err))
		return
	}
	commen.OKWithMsg("修改密码成功", c)
}
