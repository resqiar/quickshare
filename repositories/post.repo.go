package repositories

import (
	"context"
	"quickshare/entities"
	"quickshare/inputs"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepoImpl struct {
	db *pgxpool.Pool
}

func InitPostRepo(db *pgxpool.Pool) PostRepo {
	return &PostRepoImpl{
		db: db,
	}
}

type PostRepo interface {
	CreatePost(input *inputs.CreatePostInput, userId string) (string, error)
	FindByID(id string) (*entities.Post, error)
}

func (repo *PostRepoImpl) CreatePost(input *inputs.CreatePostInput, userId string) (string, error) {
	var target entities.Post

	SQL := "INSERT INTO posts(title, content, cover_url, author_id) VALUES ($1, $2, $3, $4) RETURNING id;"
	row := repo.db.QueryRow(context.Background(), SQL, input.Title, input.Content, input.CoverURL, userId)
	if err := row.Scan(&target.ID); err != nil {
		return "", err
	}

	return target.ID, nil
}

func (repo *PostRepoImpl) FindByID(id string) (*entities.Post, error) {
	var target entities.Post

	SQL := "SELECT id, title, content, author_id FROM posts WHERE id = $1;"
	row := repo.db.QueryRow(context.Background(), SQL, id)
	if err := row.Scan(
		&target.ID,
		&target.Title,
		&target.Content,
		&target.AuthorID,
	); err != nil {
		return nil, err
	}

	return &target, nil
}
