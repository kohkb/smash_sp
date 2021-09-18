package repository

import (
	"github.com/kohkb/smash_sp/pkg/domain"
	"github.com/kohkb/smash_sp/pkg/infra"
)

var db, _ = infra.NewDbHandler()

type FavoriteRepository interface {
	FindAll() (*[]domain.Favorite, error)
	Create() (*domain.Favorite, error)
}

type FavoriteRepositoryImpl struct{}

func NewFavoriteRepositoryImpl() (*FavoriteRepositoryImpl, error) {
	return &FavoriteRepositoryImpl{}, nil
}

func (f *FavoriteRepositoryImpl) FindAll() (*[]domain.Favorite, error) {
	var favorites []domain.Favorite
	if err := db.Find(&favorites).Error; err != nil {
		return nil, err
	}
	return &favorites, nil
}

func (f *FavoriteRepositoryImpl) Create(videoId string) (*domain.Favorite, error) {
	favorite := domain.Favorite{VideoId: videoId}

	if err := db.Create(&favorite).Error; err != nil {
		return nil, err
	}

	return &favorite, nil
}

func (f *FavoriteRepositoryImpl) Delete(videoId string) (*domain.Favorite, error) {
	favorite := domain.Favorite{VideoId: videoId}

	if err := db.Delete(&favorite).Error; err != nil {
		return nil, err
	}

	return &favorite, nil
}
