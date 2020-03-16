package categoryrepository

import "github.com/DACNPM-Group5/leesin/pkg/model"

type Repository interface {
	Create(category *model.Category) error
	Get(id uint) (model.Category, error)
	Gets() ([]model.Category, error)
}

type repository struct {
	dbRepo DBRepository
}

func (r *repository) Create(category *model.Category) error {
	return r.dbRepo.Create(category)
}

func (r *repository) Get(id uint) (model.Category, error) {
	return r.dbRepo.Get(id)
}

func (r *repository) Gets() ([]model.Category, error) {
	return r.dbRepo.Gets()
}

func NewRepository(dbRepo DBRepository) Repository {
	return &repository{dbRepo: dbRepo}
}
