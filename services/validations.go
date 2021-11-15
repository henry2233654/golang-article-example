package services

import (
	"fmt"
	"golang-article-example/repositories"
)

type ValidateFunc func(errCh chan<- error)

func Validate(validations ...ValidateFunc) error {
	errors := make([]*ServiceErrorDetail, 0)
	errCh := make(chan error, len(validations))
	for _, validation := range validations {
		validation(errCh)
		err := <-errCh
		switch err := err.(type) {
		case nil:
			continue
		case *ServiceErrorDetail:
			errors = append(errors, err)
		default:
			return err
		}
	}
	if len(errors) > 0 {
		return NewValidateError(errors)
	}
	return nil
}

func NotExistMessage(resourceName string, expected interface{}) string {
	return fmt.Sprintf("no [%s] was found matching the [%v]", resourceName, expected)
}

func BeingUsedMessage(resourceName string) string {
	return fmt.Sprintf("this [%s] is being used", resourceName)
}

func ValidateArticleIdExist(repo repositories.IArticle, index *int, fieldName string, id uint) ValidateFunc {
	return func(errCh chan<- error) {
		var err error
		defer func() { errCh <- err }()
		isExist, _, err := repo.Get(id)
		if err != nil {
			return
		}
		if !isExist {
			err = NewServiceErrorDetail(index, &fieldName, NotExist, NotExistMessage("article", id), nil)
			return
		}
	}
}
