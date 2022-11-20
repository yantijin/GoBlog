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
	regUserInfo, err := service.AllServiceApp.UserService.RegisterUser(commen.GVA_DB, ruser)
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
	loginUser, err := service.AllServiceApp.UserService.LogInUser(commen.GVA_DB, &ruser)
	if err != nil {
		commen.GVA_LOG.Error("登录失败", zap.Error(err))
		commen.FailedWithDetailed(loginUser, err.Error(), c)
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
		commen.FailedWithMsg("登录失败", c)
		return
	}
	commen.OkWithDetailed(model.LogInUserResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
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
	_, err := service.AllServiceApp.UserService.ChangePwd(commen.GVA_DB, user, cp.NewPassword)
	if err != nil {
		commen.FailedWithMsg("修改密码失败,原密码与当前账户不符", c)
		commen.GVA_LOG.Error("修改密码失败!", zap.Error(err))
		return
	}
	commen.OKWithMsg("修改密码成功", c)
}

// 编辑用户信息，首先需要获取，需要首先jwt验证id
func (u *UserController) GetEditUserInfo(c *gin.Context) {
	userId := utils.GetUserID(c)
	user, err := service.AllServiceApp.FindUser(commen.GVA_DB, userId)
	if err != nil {
		commen.GVA_LOG.Error("验证用户失败，请登录", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	usp := model.UserResponse{
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
	commen.OkWithDetailed(&usp, "获取用户信息成功", c)
}

// 对用户的信息进行更改，需要首先经过jwt的验证
func (u *UserController) PostEditUserInfo(c *gin.Context) {
	userId := utils.GetUserID(c)
	var asp model.UserResponse
	err := c.ShouldBindJSON(&asp)
	if err != nil {
		commen.GVA_LOG.Error("绑定数据失败，请检查", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}

	// 信息入库
	err = service.AllServiceApp.UserService.UpdateColumns(commen.GVA_DB, userId, map[string]interface{}{
		"username": asp.UserName,
		"email":    asp.Email,
		"nickname": asp.NickName,
		"avatar":   asp.Avatar,
	})
	if err != nil {
		commen.GVA_LOG.Error("更新信息失败，请检查！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	commen.OKWithMsg("更新信息成功", c)
}
