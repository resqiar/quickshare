CREATE TABLE posts (
    id TEXT PRIMARY KEY DEFAULT uuid_generate_v4(),

    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    expired_at TIMESTAMP,

    title VARCHAR(100) NOT NULL,
    content TEXT DEFAULT '',
    cover_url TEXT DEFAULT '',

    author_id TEXT REFERENCES users(id)
);
