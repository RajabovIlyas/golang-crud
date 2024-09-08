-- +goose Up


CREATE TYPE formats AS ENUM ('video', 'photo', 'music', 'other');

create table files
(
    id         uuid primary key,
    file_name  text unique not null,
    path       text        not null,
    format     formats     not null,
    size       int8        not null,
    user_id    uuid        references users (id) on delete cascade,
    created_at TIMESTAMP   not null,
    updated_at TIMESTAMP   not null
);

-- +goose Down

DROP TABLE files;

DROP TYPE formats;