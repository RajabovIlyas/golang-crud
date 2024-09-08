-- +goose Up
alter table users rename column name to username;
alter table users add column password varchar(255) not null;

-- +goose Down

alter table users drop column password;
alter table users rename column username to name;