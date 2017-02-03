
-- +goose Up

CREATE TABLE posts (
       id serial UNIQUE NOT NULL,
       user_id int,
       url TEXT,
       content TEXT,
       title TEXT NOT NULL,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       deleted_at TIMESTAMP,

       PRIMARY KEY (id)
);

-- +goose Down

DROP TABLE posts;
