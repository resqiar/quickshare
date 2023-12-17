package repositories

import (
	"context"
	"quickshare/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepoImpl struct {
	db *pgxpool.Pool
}

func InitUserRepo(db *pgxpool.Pool) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}

type UserRepo interface {
	CreateUser(*entities.User) (string, error)
	FindByEmail(email string) (*entities.User, error)
	FindByID(id string) (*entities.User, error)
}

func (repo *UserRepoImpl) CreateUser(user *entities.User) (string, error) {
	var target entities.User

	SQL := "INSERT INTO users(username, email, bio, picture_url) VALUES ($1, $2, $3, $4) RETURNING id;"
	row := repo.db.QueryRow(context.Background(), SQL, user.Username, user.Email, "", user.PictureURL)
	if err := row.Scan(&target.ID); err != nil {
		return "", err
	}

	return target.ID, nil
}

func (repo *UserRepoImpl) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	SQL := "SELECT * FROM users WHERE email = $1"
	row := repo.db.QueryRow(context.Background(), SQL, email)
	if err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Username,
		&user.Email,
		&user.Bio,
		&user.PictureURL,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepoImpl) FindByID(id string) (*entities.User, error) {
	var user entities.User

	SQL := "SELECT * FROM users WHERE id = $1"
	row := repo.db.QueryRow(context.Background(), SQL, id)
	if err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Username,
		&user.Email,
		&user.Bio,
		&user.PictureURL,
	); err != nil {
		return nil, err
	}

	return &user, nil
}
