package articlerepository

import "github.com/DACNPM-Group5/leesin/pkg/model"

type Repository interface {
	Create(a *model.Article) error
	Get(id uint) (model.Article, error)
	Gets(page, perPage int) ([]model.Article, error)
}

type repository struct {
	dbRepo DBRepository
}

func (r *repository) Create(a *model.Article) error {
	return r.dbRepo.Create(a)
}

func (r *repository) Get(id uint) (model.Article, error) {
	return r.dbRepo.Get(id)
}

func (r *repository) Gets(page, perPage int) ([]model.Article, error) {
	return r.dbRepo.Gets(page-1, perPage)
}

func NewRepository(dbRepo DBRepository) Repository {
	return &repository{dbRepo: dbRepo}
}
