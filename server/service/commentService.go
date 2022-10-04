package service

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/utils"

	"gorm.io/gorm"
)

// 主要是comment的CRUD
type CommentService struct{}

func (cs *CommentService) CreateComment(db *gorm.DB, ct *model.Comments) error {
	return commen.GVA_DB.Create(ct).Error
}

func (cs *CommentService) UpdateComment(db *gorm.DB, ct *model.Comments) error {
	return commen.GVA_DB.Save(ct).Error
}

func (cs *CommentService) FindComment(db *gorm.DB, commentId int64) (model.Comments, error) {
	comment := model.Comments{}
	err := utils.NewSqlCnd().Where("id=?", commentId).FindOne(commen.GVA_DB, &comment)
	return comment, err
}

func (cs *CommentService) DelComment(db *gorm.DB, commentId int64) error {
	err := commen.GVA_DB.Delete("id=?", commentId).Error
	return err
}

func (cs *CommentService) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Comments{}).Where("id=?", id).UpdateColumn(name, value).Error
	return err
}

func (cs *CommentService) UpdateColumns(db *gorm.DB, id int64, columns map[string]interface{}) error {
	err := db.Model(&model.Comments{}).Where("id=?").Updates(columns).Error
	return err
}

// 发表评论,article以及comment后续都需要添加回复数量+1属性
func (cs *CommentService) PublishComment(ct *model.Comments) error {
	err := commen.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 首先对comment进行入库,
		err := cs.CreateComment(tx, ct)
		if err != nil {
			return err
		}
		// 然后对对应实体的comment_count以及last_comment_time进行更新
		if ct.EntityType == model.EntityArticle {
			err = AllServiceApp.OnCommentArticle(tx, ct)
			if err != nil {
				return err
			}
		} else if ct.EntityType == model.EntityComment {
			err = cs.OnCommentComment(tx, ct)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

//对comment中的commnet_count以及idx_last_comment_time进行更新
func (cs *CommentService) OnCommentComment(db *gorm.DB, ct *model.Comments) error {
	err := cs.UpdateColumns(db, ct.ID, map[string]interface{}{
		"comment_count":         gorm.Expr("comment_count + 1"),
		"idx_last_comment_time": ct.CreatedAt.Unix(),
	})
	if err != nil {
		return err
	}
	return nil
}

// 获取所有某个entity中某个ID对应的所有comments,TODO: 不要一次性全拉取,分页分批拉取
func (cs *CommentService) GetComments(db *gorm.DB, entityType string, entityId int64) (comments []model.Comments) {
	if entityType == model.EntityArticle {
		db.Model(&model.Comments{}).Where("entity_id=?", entityId).Order("id DESC").Find(&comments)
	} else if entityType == model.EntityComment {
		db.Model(&model.Comments{}).Where("entity_id=?", entityId).Order("id DESC").Find(&comments)
	}
	return
}
