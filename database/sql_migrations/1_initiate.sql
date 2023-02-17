-- +migrate Up
-- +migrate StatementBegin

create table person (
    id serial primary key not null ,
    first_name varchar(256),
    last_name varchar(256)
)

-- +migrate StatementEnd