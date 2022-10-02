package service

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/utils"

	"gorm.io/gorm"
)

// 主要提供文章的CRUD
type ArticleService struct{}

func (as *ArticleService) CreateArticle(at *model.Article) error {
	err := commen.GVA_DB.Create(at).Error
	return err
}

func (as *ArticleService) UpdateArticle(at *model.Article) error {
	err := commen.GVA_DB.Save(at).Error
	return err
}

func (as *ArticleService) FindArticle(id int64) (*model.Article, error) {
	res := model.Article{}
	err := utils.NewSqlCnd().Where("id=?", id).FindOne(commen.GVA_DB, &res)
	return &res, err
}

func (as *ArticleService) DelArtilce(id int64) error {
	err := commen.GVA_DB.Delete(&model.Article{}, "id=?", id).Error
	return err
}

func (as *ArticleService) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) error {
	return db.Model(&model.Article{}).Where("id=?", id).UpdateColumn(name, value).Error
}

// 对浏览量+1
func (as *ArticleService) PlusViewCount(articleId int64) error {
	return as.UpdateColumn(commen.GVA_DB, articleId, "view_count", gorm.Expr("view_count + 1"))
}
