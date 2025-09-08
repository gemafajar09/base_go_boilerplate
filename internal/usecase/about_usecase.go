package usecase

import (
	"go-project/internal/domain"
	"go-project/internal/repository"
	"log"
)

type AboutUsecase struct {
	aboutRepository repository.AboutRepository
}

func NewAboutUsecase(aboutRepository repository.AboutRepository) *AboutUsecase {
	return &AboutUsecase{
		aboutRepository: aboutRepository,
	}
}

func (uc *AboutUsecase) GetAbout() ([]domain.About, error) {
	return uc.aboutRepository.GetAbout()
}

func (uc *AboutUsecase) CreateAbout(about domain.About) (domain.About, error) {
	return uc.aboutRepository.CreateAbout(about)
}

func (uc *AboutUsecase) EditAbout(id int) (domain.About, error) {
	return uc.aboutRepository.EditAbout(id)
}

func (uc *AboutUsecase) UpdateAbout(about domain.About) (domain.About, error) {
	log.Println(about)
	return uc.aboutRepository.UpdateAbout(about)
}
