package repositories

import (
	"context"
	"quickshare/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TileRepo interface {
	GetAll() ([]entities.Tile, error)
}

type TileRepoImpl struct {
	db *pgxpool.Pool
}

func InitTileRepo(db *pgxpool.Pool) TileRepo {
	return &TileRepoImpl{
		db: db,
	}
}

func (r *TileRepoImpl) GetAll() ([]entities.Tile, error) {
	var tiles []entities.Tile

	SQL := "SELECT * FROM tiles;"
	rows, err := r.db.Query(context.Background(), SQL)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tile entities.Tile

		if err = rows.Scan(
			&tile.ID,
			&tile.Name,
			&tile.Description,
			&tile.Color,
			&tile.Price,
			&tile.Width,
			&tile.Length,
			&tile.Height,
			&tile.Weight,
			&tile.Quantity,
			&tile.CreatedAt,
			&tile.UpdatedAt,
		); err != nil {
			return nil, err
		}

		tiles = append(tiles, tile)
	}

	return tiles, nil
}
