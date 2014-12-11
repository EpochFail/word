
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
alter table votes
add ipaddress character varying;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
alter table votes
drop ipaddress;
