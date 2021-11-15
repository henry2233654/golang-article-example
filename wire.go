//+build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"golang-article-example/repositories"
	"golang-article-example/services"
	"golang-article-example/services/contexts"
	"golang-article-example/web/controllers"
	"gorm.io/gorm"
)

func InitApp(
	db *gorm.DB,
	webEngine *gin.Engine,
) *App {
	wire.Build(
		wire.Struct(new(App), "*"),
		wire.Struct(new(Web), "*"),
		wire.Value(repositories.ArticleFactory(repositories.NewArticle)),
		wire.Struct(new(contexts.ArticleFactory), "*"),
		wire.Struct(new(services.Article), "*"),
		wire.Bind(new(services.IArticle), new(*services.Article)),
		wire.Struct(new(controllers.Article), "*"),
	)
	return &App{}
}
