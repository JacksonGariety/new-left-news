
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE posts
ADD COLUMN parent_post_id integer;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE posts
DROP COLUMN parent_post_id;
