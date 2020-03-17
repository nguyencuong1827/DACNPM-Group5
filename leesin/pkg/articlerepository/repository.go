package articlerepository

import (
	"encoding/json"
	"github.com/DACNPM-Group5/leesin/pkg/model"
)

type Repository interface {
	Create(a *model.Article) error
	Get(id uint) (model.Article, error)
	Gets(page, perPage int) ([]model.Article, error)
}

type repository struct {
	dbRepo DBRepository
}

func (r *repository) Create(a *model.Article) error {
	err := serializeArticle(a)
	if err != nil {
		return err
	}

	return r.dbRepo.Create(a)
}

func (r *repository) Get(id uint) (article model.Article, err error) {
	article, err = r.dbRepo.Get(id)
	if err != nil {
		return
	}

	err = unmarshalArticle(&article)
	return
}

func (r *repository) Gets(page, perPage int) (
	articles []model.Article, err error) {
	articles, err = r.dbRepo.Gets(page-1, perPage)
	if err != nil {
		return
	}

	for _, article := range articles {
		err = unmarshalArticle(&article)
		if err != nil {
			return
		}
	}

	return articles, nil
}

func NewRepository(dbRepo DBRepository) Repository {
	return &repository{dbRepo: dbRepo}
}

func serializeArticle(a *model.Article) error {
	bodyBytes, err := json.Marshal(a)
	if err != nil {
		return err
	}

	a.BodyStr = string(bodyBytes)
	return nil
}

func unmarshalArticle(a *model.Article) error {
	return json.Unmarshal([]byte(a.BodyStr), &a.Body)
}
