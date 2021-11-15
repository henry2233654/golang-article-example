package main

import (
	"github.com/gin-gonic/gin"
	"golang-article-example/migrations"
	_webControllers "golang-article-example/web/controllers"
	"gorm.io/gorm"
	"reflect"
)

type App struct {
	Web
	DB        *gorm.DB
	WebEngine *gin.Engine
}

func (a *App) Serve() error {
	a.migrate()
	a.Web.register(a.WebEngine)
	return a.WebEngine.Run(":8081")
}

func (a *App) migrate() {
	migrations.Migrate(a.DB)
}

type Web struct {
	Article *_webControllers.Article
}

func (w Web) register(e *gin.Engine) {
	type controller interface {
		Route(*gin.Engine)
	}
	v := reflect.ValueOf(w)
	for i := 0; i < v.NumField(); i++ {
		c, ok := v.Field(i).Interface().(controller)
		if !ok {
			continue
		}
		c.Route(e)
	}
}
