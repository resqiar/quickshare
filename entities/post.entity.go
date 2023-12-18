package entities

import "time"

type Post struct {
	ID string

	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt time.Time

	Title    string
	Content  string
	CoverURL string

	AuthorID string
}

type JoinPost struct {
	ID string

	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt time.Time

	Title    string
	Content  string
	CoverURL string

	AuthorID string
	Username string
}
