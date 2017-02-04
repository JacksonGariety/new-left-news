
-- +goose Up

CREATE TABLE points (
       id serial UNIQUE NOT NULL,
       user_id int,
       post_id int,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       deleted_at TIMESTAMP,

       PRIMARY KEY (id)
);

-- +goose Down

DROP TABLE points;
