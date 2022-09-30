package service

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/utils"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 用户注册，登录，修改密码等协议

type UserService struct{}

// 注册服务
func (us *UserService) RegisterUser(u model.User) (ru model.User, err error) {
	// 首先查询用户名是否已经注册
	var user model.User
	if !errors.Is(utils.NewSqlCnd().Where("username=?", u.UserName).FindOne(commen.GVA_DB, &user), gorm.ErrRecordNotFound) {
		return ru, errors.New("用户名已注册，请更改")
	}
	// 将password进行编码保存,并创建唯一标识符
	u.Password, err = utils.GetHashPwd(u.Password)
	if err != nil {
		return ru, err
	}
	u.UUID = uuid.NewV4()
	err = commen.GVA_DB.Create(&u).Error
	return u, err
}

// 登录服务,校验用户名,密码
func (us *UserService) LogInUser(u *model.User) (ru *model.User, err error) {
	if commen.GVA_DB == nil {
		return nil, fmt.Errorf("数据库不存在,请检查")
	}
	var user model.User
	// 确认用户名存在
	err = utils.NewSqlCnd().Where("username=?", u.UserName).FindOne(commen.GVA_DB, &user)
	// 查看密码是够能对应
	if err == nil {
		if ok := utils.CheckPwd(u.Password, user.Password); !ok {
			return nil, fmt.Errorf("密码错误,请检查")
		}
		commen.GVA_LOG.Info("校验通过,用户登录")
		return &user, nil
	}
	return &user, err
}

// 修改密码
func (us *UserService) ChangePwd(u *model.User, newPwd string) (ru *model.User, err error) {
	var user model.User
	err = utils.NewSqlCnd().Where("ID=?", u.ID).FindOne(commen.GVA_DB, &user)
	if err == nil {
		// 校验原来的密码是对的
		if ok := utils.CheckPwd(u.Password, user.Password); !ok {
			return nil, fmt.Errorf("原密码错误")
		}
		// 校验通过,对密码进行修改
		user.Password = newPwd
		commen.GVA_DB.Save(&user)
		return &user, nil
	}
	return nil, err
}

// 删除用户
func (us *UserService) DelUser(id int) (err error) {
	var user model.User
	err = commen.GVA_DB.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
