// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"golang-article-example/repositories"
	"golang-article-example/services"
	"golang-article-example/services/contexts"
	"golang-article-example/web/controllers"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitApp(db *gorm.DB, webEngine *gin.Engine) *App {
	article := &services.Article{}
	articleFactory := _wireArticleFactoryValue
	contextsArticleFactory := &contexts.ArticleFactory{
		DB:                 db,
		ArticleRepoFactory: articleFactory,
	}
	controllersArticle := &controllers.Article{
		Service:    article,
		CtxFactory: contextsArticleFactory,
	}
	web := Web{
		Article: controllersArticle,
	}
	app := &App{
		Web:       web,
		DB:        db,
		WebEngine: webEngine,
	}
	return app
}

var (
	_wireArticleFactoryValue = repositories.ArticleFactory(repositories.NewArticle)
)
