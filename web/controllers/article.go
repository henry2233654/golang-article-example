package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-article-example/services"
	"golang-article-example/services/contexts"
	"golang-article-example/web/cores"
	"golang-article-example/web/cores/di"
	"golang-article-example/web/inputmodels"
	"golang-article-example/web/viewmodels"
	"net/http"
)

type Article struct {
	Service    services.IArticle
	CtxFactory *contexts.ArticleFactory
}

func (ctr *Article) Route(e *gin.Engine) {
	article := e.Group("/articles")
	{
		article.GET("", di.AutoDi(ctr.List)...)
		article.POST("", di.AutoDi(ctr.Create)...)
		article.PUT("", di.AutoDi(ctr.Update)...)
		article.DELETE("", di.AutoDi(ctr.Delete)...)
		article.GET("/:id", di.AutoDi(ctr.Get)...)
	}
	//e.GET("/import/article", di.AutoWsDi(ctr.Import))
}

// List
// @Summary List Articles
// @Description List articles by list parameters
// @ID list-articles
// @Tags Article
// @Param X-USER-JWT header string true "User Info JWT"
// @Param ListParam query inputmodels.ListParam true "ListParam"
// @Success 200 {object} []viewmodels.Article "Success."
// @Failure 500 {object} viewmodels.Error "Internal error."
// @Router /articles [get]
func (ctr *Article) List(listParam inputmodels.ListParam) (*cores.Response, error) {
	ctx := ctr.CtxFactory.NewContext(nil)
	articles, err := ctr.Service.ListArticles(ctx)
	if err != nil {
		return useDefaultErrorHandler(err)
	}
	respBody := make([]viewmodels.Article, len(articles))
	for i, article := range articles {
		respBody[i] = viewmodels.Article{
			ID:        article.ID,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
			No:        article.No,
			Name:      article.Name,
			QuickCode: article.QuickCode,
		}
	}
	resp := cores.NewResponse(http.StatusOK, respBody)
	return resp, nil
}

// Create
// @Summary Create Article
// @Description Create a batch of articles
// @ID create-articles
// @Tags Article
// @Param X-USER-JWT header string true "User Info JWT"
// @Param Item body []inputmodels.CreateArticle true "CreateArticles"
// @Success 201 "Success."
// @Success 400 {object} viewmodels.Error "Bad Request."
// @Success 409 {object} viewmodels.Error "Conflict."
// @Failure 500 {object} viewmodels.Error "Internal error."
// @Router /articles [post]
func (ctr *Article) Create(createArticles []inputmodels.CreateArticle) (*cores.Response, error) {
	dtos := make([]services.CreateArticle, len(createArticles))
	for i, createArticle := range createArticles {
		dtos[i] = services.CreateArticle{
			No:        createArticle.No,
			Name:      createArticle.Name,
			QuickCode: createArticle.QuickCode,
		}
	}
	ctx := ctr.CtxFactory.NewContext(nil)
	if err := ctr.Service.BatchCreateArticles(ctx, dtos); err != nil {
		return useDefaultErrorHandler(err)
	}
	return cores.NewResponse(http.StatusCreated, nil), nil
}

// Get
// @Summary Get Article
// @Description Get article by ID
// @ID get-article
// @Tags Article
// @Param X-USER-JWT header string true "User Info JWT"
// @Param id path string true "ID"
// @Success 200 {object} viewmodels.Article "Success."
// @Success 404 {object} viewmodels.Error "Not found."
// @Failure 500 {object} viewmodels.Error "Internal error."
// @Router /articles/{id} [get]
func (ctr *Article) Get(id uint) (*cores.Response, error) {
	ctx := ctr.CtxFactory.NewContext(nil)
	article, err := ctr.Service.RetrieveArticle(ctx, id)
	if err != nil {
		return useDefaultErrorHandler(err)
	}
	respBody := viewmodels.Article{
		ID:        article.ID,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		No:        article.No,
		Name:      article.Name,
		QuickCode: article.QuickCode,
	}
	return cores.NewResponse(http.StatusOK, respBody), nil
}

// Update
// @Summary Update Article
// @Description Update a batch of articles
// @ID update-article
// @Tags Article
// @Param X-USER-JWT header string true "User Info JWT"
// @Param Article body []inputmodels.UpdateArticle true "Update parameters."
// @Success 204 "Success."
// @Success 404 {object} viewmodels.Error "Not found."
// @Success 409 {object} viewmodels.Error "Conflict."
// @Failure 500 {object} viewmodels.Error "Internal error."
// @Router /articles [put]
func (ctr *Article) Update(updateArticles []inputmodels.UpdateArticle) (*cores.Response, error) {
	dtos := make([]services.UpdateArticle, len(updateArticles))
	for i, updateArticle := range updateArticles {
		dtos[i] = services.UpdateArticle{
			ID:        updateArticle.ID,
			No:        updateArticle.No,
			Name:      updateArticle.Name,
			QuickCode: updateArticle.QuickCode,
		}
	}
	ctx := ctr.CtxFactory.NewContext(nil)
	if err := ctr.Service.BatchUpdateArticles(ctx, dtos); err != nil {
		return useDefaultErrorHandler(err)
	}
	return cores.NewResponse(http.StatusNoContent, nil), nil
}

// Delete
// @Summary Delete Article
// @Description Delete a batch of articles
// @ID delete-article
// @Tags Article
// @Param X-USER-JWT header string true "User Info JWT"
// @Param id path string true "ID"
// @Success 204 "Success."
// @Success 400 {object} viewmodels.Error "Bad Request."
// @Success 404 {object} viewmodels.Error "Not found."
// @Failure 500 {object} viewmodels.Error "Internal error."
// @Router /articles [Delete]
func (ctr *Article) Delete(ids []uint) (*cores.Response, error) {
	ctx := ctr.CtxFactory.NewContext(nil)
	err := ctr.Service.BatchDeleteArticles(ctx, ids)
	if err != nil {
		return useDefaultErrorHandler(err)
	}
	return cores.NewResponse(http.StatusNoContent, nil), nil
}

// Import
// @Summary [Websocket] Import Article
// @Description Import article
// @ID import-article
// @Tags Import
// @Response 101 {string} string "Switching To Websocket"
// @Router /import/article [get]
//func (ctr *Article) Import(c context.Context, inCh <-chan inputmodels.CreateArticle, respCh chan<- *cores.WsResponse) {
//	for i := 0; ; i++ {
//		select {
//		case createArticle := <-inCh:
//			var dto services.CreateArticle
//			if err := ctr.Converter.Convert(createArticle, &dto); err != nil {
//				respCh <- cores.NewErrorWsResponse(err)
//				break
//			}
//			ctx := ctr.CtxFactory.NewContext(nil)
//			article, err := ctr.Service.CreateArticle(ctx, i, dto)
//			if err != nil {
//				respCh <- cores.NewErrorWsResponse(err)
//				break
//			}
//			respCh <- cores.NewImportedWsResponse(article.ID)
//		case <-c.Done():
//			return
//		}
//	}
//}
