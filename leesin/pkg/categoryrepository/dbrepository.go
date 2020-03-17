package categoryrepository

import (
	"DACNPM-Group5/leesin/pkg/model"
	"github.com/jinzhu/gorm"
)

type DBRepository interface {
	Create(category *model.Category) error
	Get(id uint) (model.Category, error)
	Gets() ([]model.Category, error)
}

type dbRepository struct {
	db *gorm.DB
}

func (d *dbRepository) Create(category *model.Category) error {
	return d.db.Create(category).Error
}

func (d *dbRepository) Get(id uint) (category model.Category, err error) {
	err = d.db.First(&category, "id = ?", id).Error
	return
}

func (d *dbRepository) Gets() (categories []model.Category, err error) {
	err = d.db.Find(&categories).Error
	return
}

func NewDBRepository(db *gorm.DB) DBRepository {
	return &dbRepository{db}
}
