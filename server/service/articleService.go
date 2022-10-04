package service

import (
	"GoLog/model"
	"GoLog/utils"

	"gorm.io/gorm"
)

// 主要提供文章的CRUD
type ArticleService struct{}

func (as *ArticleService) CreateArticle(db *gorm.DB, at *model.Article) error {
	err := db.Create(at).Error
	return err
}

func (as *ArticleService) UpdateArticle(db *gorm.DB, at *model.Article) error {
	err := db.Save(at).Error
	return err
}

func (as *ArticleService) FindArticle(db *gorm.DB, id int64) (*model.Article, error) {
	res := model.Article{}
	err := utils.NewSqlCnd().Where("id=?", id).FindOne(db, &res)
	return &res, err
}

func (as *AllService) FindArticles(db *gorm.DB, cnd *utils.SqlCnd) (articleList []model.Article) {
	cnd.Find(db, &articleList)
	return
}

func (as *ArticleService) DelArtilce(db *gorm.DB, id int64) error {
	err := db.Delete(&model.Article{}, "id=?", id).Error
	return err
}

func (as *ArticleService) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) error {
	return db.Model(&model.Article{}).Where("id=?", id).UpdateColumn(name, value).Error
}

// 对浏览量+1
func (as *ArticleService) PlusViewCount(db *gorm.DB, articleId int64) error {
	return as.UpdateColumn(db, articleId, "view_count", gorm.Expr("view_count + 1"))
}

// 获取文章列表,从最近的开始，这里需要改进，如果数据库的表很大，查询会比较慢; 可以添加limit和cursor，一点点查
//另外，如果添加私有文章功能，则需要再加一个条件判断
func (as *ArticleService) GetUserArticles(db *gorm.DB, userId int64) (articleList []model.Article) {
	utils.NewSqlCnd().Eq("user_id", userId).Desc("id").Find(db, articleList)
	return articleList
}

func (as *ArticleService) PlusCommentCount(db *gorm.DB, articleId int64) error {
	return as.UpdateColumn(db, articleId, "comment_count", gorm.Expr("comment_count+1"))
}
