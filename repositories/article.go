package repositories

import (
	"golang-article-example/entities"

	"gorm.io/gorm"
)

type ArticleFactory func(ctx *GormDBContext) IArticle

type IArticle interface {
	ListArticles() (articles []*entities.Article, err error)
	Save(article *entities.Article) (err error)
	Get(id uint) (isExist bool, article *entities.Article, err error)
	Delete(article *entities.Article) (err error)
}

type Article struct {
	GormRepository
}

func NewArticle(ctx *GormDBContext) IArticle {
	repository := new(Article)
	repository.SetDBContext(ctx)
	return repository
}

func (repo *Article) Save(article *entities.Article) (err error) {
	err = repo.DB().Save(article).Error
	if err != nil && isCausedByUniqueConstraint(err) {
		err = NewUniqueConstrainError(err)
	}
	return
}

func (repo *Article) Delete(article *entities.Article) (err error) {
	err = repo.DB().Delete(article).Error
	return
}

func (repo *Article) ListArticles() (articles []*entities.Article, err error) {
	var tempArticles []*entities.Article
	err = repo.DB().Order("id").Find(&tempArticles).Error
	articles = tempArticles
	return
}

func (repo *Article) Get(id uint) (isExist bool, article *entities.Article, err error) {
	var tempProvider entities.Article
	err = repo.DB().First(&tempProvider, id).Error
	if err != nil {
		isExist = false
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	isExist = true
	article = &tempProvider
	return
}
