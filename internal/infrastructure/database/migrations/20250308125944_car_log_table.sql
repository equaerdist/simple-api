-- +goose Up
-- +goose StatementBegin
create table car_log (
    id bigserial primary key,
    car_id bigint not null,
    model_name varchar(50) not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table car_log;
-- +goose StatementEnd
