
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE points
ADD COLUMN vote integer;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE points
DROP COLUMN vote;
-- +goose Up -- SQL in section 'Up' is executed when this migration is applied -- +goose Down -- SQL section 'Down' is executed when this migration is rolled back
