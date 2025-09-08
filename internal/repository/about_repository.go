package repository

import (
	"go-project/internal/domain"

	"gorm.io/gorm"
)

type AboutRepository interface {
	GetAbout() ([]domain.About, error)
	CreateAbout(about domain.About) (domain.About, error)
	EditAbout(id int) (domain.About, error)
	UpdateAbout(about domain.About) (domain.About, error)
}

type aboutRepository struct {
	db *gorm.DB
}

func NewAboutRepository(db *gorm.DB) AboutRepository {
	return &aboutRepository{db: db}
}

func (r *aboutRepository) GetAbout() ([]domain.About, error) {
	var abouts []domain.About
	err := r.db.Find(&abouts).Error
	return abouts, err
}

func (r *aboutRepository) CreateAbout(about domain.About) (domain.About, error) {
	err := r.db.Create(&about).Error
	return about, err
}

func (r *aboutRepository) EditAbout(id int) (domain.About, error) {
	var about domain.About
	err := r.db.Where("id = ?", id).Find(&about).Error
	return about, err
}

func (r *aboutRepository) UpdateAbout(about domain.About) (domain.About, error) {
	err := r.db.Model(&domain.About{}).
		Where("id = ?", about.Id).
		Updates(domain.About{
			Title:   about.Title,
			Content: about.Content,
		}).Error
	return about, err
}
