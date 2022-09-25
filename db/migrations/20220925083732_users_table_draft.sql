-- +goose Up
-- +goose StatementBegin
create table user(
    id bigint not null,
    username text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
-- +goose StatementEnd
