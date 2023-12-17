package services

import (
	"quickshare/entities"
	"quickshare/repositories"
)

type TileService interface {
	GetAllTiles() ([]entities.Tile, error)
}

type TileServiceImpl struct {
	Repository repositories.TileRepo
}

func (s *TileServiceImpl) GetAllTiles() ([]entities.Tile, error) {
	return s.Repository.GetAll()
}
