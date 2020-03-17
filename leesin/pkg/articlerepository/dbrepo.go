package articlerepository

import (
	"DACNPM-Group5/leesin/pkg/model"
	"github.com/jinzhu/gorm"
)

type DBRepository interface {
	Get(id uint) (model.Article, error)
	Gets(offset, limit int) ([]model.Article, error)
	Create(a *model.Article) error
}

type dbRepository struct {
	db *gorm.DB
}

func (d *dbRepository) Get(id uint) (article model.Article, err error) {
	err = d.db.First(&article, "id = ?", id).Error
	return
}

func (d *dbRepository) Gets(offset, limit int) ([]model.Article, error) {
	var result []model.Article
	err := d.db.Offset(offset).Limit(limit).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *dbRepository) Create(a *model.Article) error {
	return d.db.Create(a).Error
}

func NewDBRepository(db *gorm.DB) DBRepository {
	return &dbRepository{db}
}
