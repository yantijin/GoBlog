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
func (us *UserService) RegisterUser(db *gorm.DB, u model.User) (ru model.User, err error) {
	// 首先查询用户名是否已经注册
	var user model.User
	if !errors.Is(utils.NewSqlCnd().Where("username=?", u.UserName).FindOne(db, &user), gorm.ErrRecordNotFound) {
		return ru, errors.New("用户名已注册，请更改")
	}
	// 将password进行编码保存,并创建唯一标识符
	u.Password, err = utils.GetHashPwd(u.Password)
	if err != nil {
		return ru, err
	}
	u.UUID = uuid.NewV4()
	err = db.Create(&u).Error
	return u, err
}

// 登录服务,校验用户名,密码
func (us *UserService) LogInUser(db *gorm.DB, u *model.User) (ru *model.User, err error) {
	if db == nil {
		return nil, fmt.Errorf("数据库不存在,请检查")
	}
	var user model.User
	// 确认用户名存在
	err = utils.NewSqlCnd().Where("username=?", u.UserName).FindOne(db, &user)
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
func (us *UserService) ChangePwd(db *gorm.DB, u *model.User, newPwd string) (ru *model.User, err error) {
	var user model.User
	err = utils.NewSqlCnd().Where("ID=?", u.ID).FindOne(db, &user)
	if err == nil {
		// 校验原来的密码是对的
		if ok := utils.CheckPwd(u.Password, user.Password); !ok {
			return nil, fmt.Errorf("原密码错误")
		}
		// 校验通过,对密码进行修改
		user.Password, err = utils.GetHashPwd(newPwd)
		if err != nil {
			return nil, fmt.Errorf("对新密码编码错误，请重试")
		}
		db.Save(&user)
		return &user, nil
	}
	return nil, err
}

// 删除用户
func (us *UserService) DelUser(db *gorm.DB, id int) (err error) {
	var user model.User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// 修改用户信息
func (us *UserService) ChangeUserInfo(db *gorm.DB, u model.User) error {
	err := db.Updates(u).Error
	return err
}

// 查找用户信息
func (us *UserService) FindUser(db *gorm.DB, id int64) (*model.User, error) {
	var user model.User
	err := db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (us *UserService) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id=?", id).Update(name, value).Error
	return
}

func (us *UserService) UpdateColumns(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id=?", id).Updates(columns).Error
	return
}
