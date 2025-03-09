-- +goose Up
-- +goose StatementBegin
create table car (
    id bigserial primary key,
    model_name varchar(50) not null,
    created_at timestamptz not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table car;
-- +goose StatementEnd
