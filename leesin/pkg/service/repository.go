package service

import (
	"github.com/DACNPM-Group5/leesin/pkg/articlerepository"
	"github.com/DACNPM-Group5/leesin/pkg/categoryrepository"
	"github.com/DACNPM-Group5/leesin/pkg/model"
)

type Repository interface {
	GetCategories() ([]model.Category, error)
	GetArticle(id uint)(model.Article, error)
	GetPreArticles(page, perPage int) ([]previewArticle, error)
}

type repository struct {
	categoryRepo categoryrepository.Repository
	articleRepo  articlerepository.Repository
}

func (r *repository) GetArticle(id uint) (model.Article, error) {
	return r.articleRepo.Get(id)
}

func (r *repository) GetCategories() ([]model.Category, error) {
	return r.categoryRepo.Gets()
}

func (r *repository) GetPreArticles(page, perPage int) ([]previewArticle, error) {
	arts, err := r.articleRepo.Gets(page, perPage)
	if err != nil {
		return nil, err
	}

	return convertToPreviewArticles(arts), nil
}

func NewRepository(articleRepo articlerepository.Repository,
	category categoryrepository.Repository) Repository {
	return &repository{
		articleRepo:  articleRepo,
		categoryRepo: category,
	}
}

func convertToPreviewArticles(articles []model.Article) []previewArticle {
	var prevs []previewArticle
	for _, article := range articles {
		prevs = append(prevs, fromArticle(article))
	}

	return prevs
}
