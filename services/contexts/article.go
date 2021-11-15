package contexts

import (
	"golang-article-example/repositories"
	"gorm.io/gorm"
)

type Article struct {
	Context
	ArticleRepo repositories.IArticle
	UserID      *uint
}

type ArticleFactory struct {
	DB                 *gorm.DB
	ArticleRepoFactory repositories.ArticleFactory
}

func (f *ArticleFactory) NewContext(userID *uint) *Article {
	dbCtx := repositories.NewGormDBContext(f.DB)
	ctx := new(Article)
	ctx.ArticleRepo = f.ArticleRepoFactory(dbCtx)
	ctx.AddDBContexts(dbCtx)
	return ctx
}
