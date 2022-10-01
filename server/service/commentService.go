package service

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/utils"

	"gorm.io/gorm"
)

// 主要是comment的CRUD
type CommentService struct{}

func (cs *CommentService) CreateComment(ct *model.Comments) error {
	return commen.GVA_DB.Create(ct).Error
}

func (cs *CommentService) UpdateComment(ct *model.Comments) error {
	return commen.GVA_DB.Save(ct).Error
}

func (cs *CommentService) FindComment(commentId int64) (model.Comments, error) {
	comment := model.Comments{}
	err := utils.NewSqlCnd().Where("id=?", commentId).FindOne(commen.GVA_DB, &comment)
	return comment, err
}

func (cs *CommentService) DelComment(commentId int64) error {
	err := commen.GVA_DB.Delete("id=?", commentId).Error
	return err
}

func (cs *CommentService) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Comments{}).Where("id=?", id).UpdateColumn(name, value).Error
	return err
}
