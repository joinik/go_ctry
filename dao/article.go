package dao

import (
	"context"
	"go_ctry/model"

	"gorm.io/gorm"
)

type ArticleDao struct {
	*gorm.DB
}

func NewArticleDao(ctx context.Context) *ArticleDao {
	return &ArticleDao{NewDBClient(ctx)}
}


// Create 创建文章
func (dao *ArticleDao) Create(art *model.Article) (err error) {
	err = dao.DB.Model(&model.Article{}).Create(&art).Error
	return
}

// UpdateArtByAuthID 根据AuthID更新文章
func (dao *ArticleDao) UpdateArtByAuthID(art *model.Article, authid uint) (err error) {
	err = dao.DB.Model(&model.Article{}).Where("author_id=?", authid).Updates(&art).Error
	return
}

// GetArtByAuthID 根据AuthID查询文章
func (dao *ArticleDao) GetArtByAuthID(authid uint) (art []*model.Article, err error) {
	err = dao.DB.Model(&model.Article{}).Where("author_id=?", authid).Find(&art).Error
	return
}

// GetArtByAreaID 根据AreaID查询文章
func (dao *ArticleDao) GetArtByAreaID(areaid uint) (art []*model.Article, err error) {
	err = dao.DB.Model(&model.Article{}).Where("area_id=?", areaid).Find(&art).Error
	return
}

// GetArtContentByArtID 根据文章id查询文章信息
func (dao *ArticleDao) GetArtByArtID(artid uint) (art *model.Article, err error) {
	err = dao.DB.Model(&model.Article{}).Where("id=?", artid).
	Preload("Cate").Preload("Area").Preload("Author").First(&art).Error
	return
}

// GetArtContentByArtID 根据文章id查询文章内容
func (dao *ArticleDao) GetArtContentByArtID(artid uint) (art *model.ArtContent, err error) {
	err = dao.DB.Model(&model.ArtContent{}).Where("article_id=?", artid).First(&art).Error
	return
}


