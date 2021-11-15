package services

import (
	"golang-article-example/entities"
	"golang-article-example/repositories"
	"golang-article-example/services/contexts"
)

type CreateArticle struct {
	No        string `json:"no"`         // 文章代號(unique)
	Name      string `json:"name"`       // 文章名稱
	QuickCode string `json:"quick_code"` // 簡碼
}

type UpdateArticle struct {
	ID        uint   `json:"id"`         // 流水號
	No        string `json:"no"`         // 文章代號(unique)
	Name      string `json:"name"`       // 文章名稱
	QuickCode string `json:"quick_code"` // 簡碼
}

type IArticle interface {
	ListArticles(c *contexts.Article) (articles []*entities.Article, err error)
	CreateArticle(c *contexts.Article, index int, createArticle CreateArticle) (article *entities.Article, err error)
	BatchCreateArticles(c *contexts.Article, createArticles []CreateArticle) (err error)
	BatchUpdateArticles(c *contexts.Article, updateArticles []UpdateArticle) (err error)
	BatchDeleteArticles(c *contexts.Article, ids []uint) (err error)
	RetrieveArticle(c *contexts.Article, id uint) (article *entities.Article, err error)
}

type Article struct {
}

func (srv *Article) ListArticles(c *contexts.Article) (articles []*entities.Article, err error) {
	return c.ArticleRepo.ListArticles()
}

func (srv *Article) RetrieveArticle(c *contexts.Article, id uint) (article *entities.Article, err error) {
	isExist, article, err := c.ArticleRepo.Get(id)
	if err != nil {
		return
	}
	if !isExist {
		err = &NotExistError{ResourceName: "article", Expected: id}
		return
	}
	return
}

func (srv *Article) CreateArticle(c *contexts.Article, index int, createArticle CreateArticle) (article *entities.Article, err error) {
	article = new(entities.Article)
	article.No = createArticle.No
	article.Name = createArticle.Name
	article.QuickCode = createArticle.QuickCode
	if err := srv.saveArticle(c, index, article); err != nil {
		return nil, err
	}
	return article, nil
}

func (srv *Article) BatchCreateArticles(c *contexts.Article, createArticles []CreateArticle) (err error) {
	fn := func() (err error) {
		for i, createArticle := range createArticles {
			if _, err := srv.CreateArticle(c, i, createArticle); err != nil {
				return err
			}
		}
		return nil
	}
	return StartTransaction(c, fn)
}

func (srv *Article) BatchUpdateArticles(c *contexts.Article, updateArticles []UpdateArticle) (err error) {
	fn := func() (err error) {
		for i, updateArticle := range updateArticles {
			id := updateArticle.ID
			article, err := srv.RetrieveArticle(c, id)
			if err != nil {
				return err
			}
			article.No = updateArticle.No
			article.Name = updateArticle.Name
			article.QuickCode = updateArticle.QuickCode
			if err := srv.saveArticle(c, i, article); err != nil {
				return err
			}
		}
		return nil
	}
	return StartTransaction(c, fn)
}

func (srv *Article) BatchDeleteArticles(c *contexts.Article, ids []uint) (err error) {
	fn := func() (err error) {
		for i, id := range ids {
			article, err := srv.RetrieveArticle(c, id)
			if err != nil {
				return err
			}
			if err = srv.validateBeforeDelete(c, i, article); err != nil {
				return err
			}
			if err = c.ArticleRepo.Delete(article); err != nil {
				return err
			}
		}
		return nil
	}
	return StartTransaction(c, fn)
}

func (srv *Article) saveArticle(c *contexts.Article, index int, article *entities.Article) error {
	if err := srv.validateBeforeSave(c, article); err != nil {
		return err
	}
	if err := c.ArticleRepo.Save(article); err != nil {
		if uniqueConstrainError, ok := err.(*repositories.UniqueConstrainError); ok {
			return NewDuplicateError(&index, uniqueConstrainError)
		}
		return err
	}
	return nil
}

func (srv *Article) validateBeforeSave(c *contexts.Article, article *entities.Article) error {
	var validations []ValidateFunc
	return Validate(validations...)
}

func (srv *Article) validateBeforeDelete(c *contexts.Article, index int, article *entities.Article) error {
	validations := []ValidateFunc{}
	return Validate(validations...)
}
