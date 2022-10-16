package service

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/utils"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 主要定义UserLike的增删改查
type UserLikeService struct{}

func (uls *UserLikeService) CreateUserLike(db *gorm.DB, ul *model.UserLike) error {
	return db.Create(ul).Error
}

func (uls *UserLikeService) FindUserLike(db *gorm.DB, id int64) (*model.UserLike, error) {
	res := model.UserLike{}
	err := utils.NewSqlCnd().Where("id=?", id).FindOne(db, &res)
	return &res, err
}

func (uls *UserLikeService) FindUserLikesByUserId(db *gorm.DB, id int64) (list []model.UserLike) {
	utils.NewSqlCnd().Where("ul_user_id=?", id).Desc("id").Find(db, &list)
	return
}

func (uls *UserLikeService) DelUserLike(db *gorm.DB, id int64) error {
	err := db.Delete(&model.UserLike{}, "id=?", id).Error
	return err
}

// 更新某个用户某一列的属性
func (uls *UserLikeService) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserLike{}).Where("id=?", id).UpdateColumn(name, value).Error
	return err
}

func (uls *UserLikeService) UpdateColumns(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserLike{}).Where("id=?", id).Updates(columns).Error
	return
}

// 统计对某个实体(文章/评论)的点赞数量
func (uls *UserLikeService) CountUserLike(db *gorm.DB, entityType string, entityId string) int64 {
	var count int64 = 0
	db.Model(&model.UserLike{}).Where("entity_type=?", entityType).Where("entity_id", entityId).Count(&count)
	return count
}

// 是否已点赞过
func (usl *UserLikeService) ExistsUserLike(db *gorm.DB, userId int64, entityType string, entityId int64) (*model.UserLike, bool) {
	var res model.UserLike
	err := utils.NewSqlCnd().Eq("ul_user_id", userId).Eq("entity_type", entityType).Eq("entity_id", entityId).FindOne(db, &res)
	return &res, err == nil
}

// 某用户是否对某些评论进行点赞过，返回点赞过的评论id
func (usl *UserLikeService) IsLiked(db *gorm.DB, userId int64, entityType string, entityIds []int64) (likedEntityIds []int64) {
	var list []model.UserLike
	utils.NewSqlCnd().Eq("ul_user_id", userId).Eq("entity_type", entityType).In("entity_id", entityIds).Find(db, list)
	for _, like := range list {
		likedEntityIds = append(likedEntityIds, like.ID)
	}
	return
}

// 点赞操作
func (usl *UserLikeService) Like(db *gorm.DB, userId int64, entityType string, entityId int64) error {
	// 首先判断是否已经点过赞
	if _, ok := usl.ExistsUserLike(db, userId, entityType, entityId); ok {
		commen.GVA_LOG.Error("已点过赞！")
		return fmt.Errorf("已点过赞")
	}
	// 调用创建服务
	return usl.CreateUserLike(db, &model.UserLike{
		UserId:     userId,
		EntityType: entityType,
		EntityId:   entityId,
	})
}

// 取消点赞操作，
func (usl *UserLikeService) UnLike(db *gorm.DB, userId int64, entityType string, entityId int64) error {
	// 首先判断是否点过赞
	ul, ok := usl.ExistsUserLike(db, userId, entityType, entityId)
	if !ok {
		commen.GVA_LOG.Error("尚未点赞！")
		return fmt.Errorf("尚未点赞")
	}
	// 调用删除点赞服务
	if err := usl.DelUserLike(db, ul.ID); err != nil {
		commen.GVA_LOG.Error("删除点赞失败", zap.Error(err))
		return err
	}
	return nil
}

// 对评论点赞，注意事务一致性
func (usl *UserLikeService) CommentLike(db *gorm.DB, userId int64, commentId int64) error {
	_, err := AllServiceApp.CommentService.FindComment(db, commentId)
	if err != nil {
		commen.GVA_LOG.Error("无法查找到对应的comment，请检查", zap.Error(err))
		return err
	}
	// 点赞之后，对应的coment点赞数量要+1,保持事务的一致性
	err = db.Transaction(func(tx *gorm.DB) error {
		// 首先进行点赞入库
		if err := usl.Like(tx, userId, model.EntityComment, commentId); err != nil {
			return err
		}
		// 对comment的like_count + 1,
		return AllServiceApp.CommentService.UpdateColumn(tx, commentId, "like_count", gorm.Expr("like_count + 1"))
	})
	if err != nil {
		return err
	}
	return nil
}

// 对文章点赞，注意事务一致性
func (usl *UserLikeService) ArticleLike(db *gorm.DB, userId int64, articleId int64) error {
	_, err := AllServiceApp.ArticleService.FindArticle(db, articleId)
	if err != nil {
		commen.GVA_LOG.Error("无法查找到对应的文章，请检查")
		return err
	}
	// 点赞之后，对应的article点赞数量+1,保持事务一致性
	err = db.Transaction(func(tx *gorm.DB) error {
		// 首先进行点赞入库
		if err := usl.Like(tx, userId, model.EntityArticle, articleId); err != nil {
			commen.GVA_LOG.Error("点赞数据入库失败，请检查", zap.Error(err))
			return err
		}
		// article中的like_count + 1
		return AllServiceApp.ArticleService.UpdateColumn(tx, articleId, "like_count", gorm.Expr("like_cout + 1"))
	})
	if err != nil {
		return err
	}
	return nil
}

// 对comment取消点赞，注意事务一致性
func (usl *UserLikeService) CommentUnLike(db *gorm.DB, userId int64, commentId int64) error {
	_, err := AllServiceApp.CommentService.FindComment(db, commentId)
	if err != nil {
		commen.GVA_LOG.Error("无法查找到对应的comment，请检查", zap.Error(err))
		return err
	}
	// 点赞之后，对应的coment点赞数量要+1,保持事务的一致性
	err = db.Transaction(func(tx *gorm.DB) error {
		// 首先进行删除点赞
		if err := usl.UnLike(tx, userId, model.EntityComment, commentId); err != nil {
			return err
		}
		// 对comment的like_count - 1,
		return AllServiceApp.CommentService.UpdateColumn(tx, commentId, "like_count", gorm.Expr("like_count - 1"))
	})
	if err != nil {
		return err
	}
	return nil
}

// 对文章取消点赞，注意事务一致性
func (usl *UserLikeService) ArticleUnLike(db *gorm.DB, userId int64, articleId int64) error {
	_, err := AllServiceApp.ArticleService.FindArticle(db, articleId)
	if err != nil {
		commen.GVA_LOG.Error("无法查找到对应的文章，请检查")
		return err
	}
	// 点赞之后，对应的article点赞数量+1,保持事务一致性
	err = db.Transaction(func(tx *gorm.DB) error {
		// 首先进行删除点赞
		if err := usl.UnLike(tx, userId, model.EntityArticle, articleId); err != nil {
			commen.GVA_LOG.Error("删除点赞失败，请检查", zap.Error(err))
			return err
		}
		// article中的like_count - 1
		return AllServiceApp.ArticleService.UpdateColumn(tx, articleId, "like_count", gorm.Expr("like_cout - 1"))
	})
	if err != nil {
		return err
	}
	return nil
}

// 获取某用户对某实体的点赞数据,返回对应的entityIds
func (usl *UserLikeService) FindPersonLikes(db *gorm.DB, userId int64, entityType string) (entityIds []int64) {
	var ullist []model.UserLike
	db.Where("ul_user_id=?", userId).Where("entity_type", entityType).Order("id DESC").Find(&ullist)
	if len(ullist) > 0 {
		for _, term := range ullist {
			entityIds = append(entityIds, term.EntityId)
		}
		return
	}
	return
}
