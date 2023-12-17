CREATE TABLE posts (
    id TEXT PRIMARY KEY DEFAULT uuid_generate_v4(),

    slug TEXT UNIQUE,

    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    expired_at TIMESTAMP,

    Title VARCHAR(100) NOT NULL,
    Content TEXT DEFAULT '',
    CoverURL TEXT DEFAULT '',

    author_id TEXT REFERENCES users(id)
);
