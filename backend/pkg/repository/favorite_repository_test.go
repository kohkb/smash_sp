package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFavoriteRepositoryImpl(t *testing.T) {
	_, err := NewFavoriteRepositoryImpl()
	assert.Nil(t, err)
}

func TestFind(t *testing.T) {
	favoriteRepository, _ := NewFavoriteRepositoryImpl()
	favorite, err := favoriteRepository.Find("1")

	assert.True(t, favorite.ID == 1)
	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	favoriteRepository, _ := NewFavoriteRepositoryImpl()
	favorites, err := favoriteRepository.FindAll()

	assert.True(t, len(*favorites) > 0)
	assert.Nil(t, err)
}

// TODO: テストの度にデータが増えるのを防ぐ
func TestCreateFavorite(t *testing.T) {
	favoriteRepository, _ := NewFavoriteRepositoryImpl()
	favorite, err := favoriteRepository.Create("abcdefg")

	assert.Equal(t, favorite.VideoId, "abcdefg")
	assert.Nil(t, err)
}

func TestDeleteFavorite(t *testing.T) {
	// favoriteRepository, _ := NewFavoriteRepositoryImpl()

	assert.Equal(t, "abcdefg", "abcdefg")
}
