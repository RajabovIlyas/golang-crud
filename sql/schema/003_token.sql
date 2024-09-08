-- +goose Up

create table tokens
(
    id               uuid primary key,
    access_token_key uuid      not null unique,
    user_id          uuid      not null references users (id) on delete cascade,
    created_at       TIMESTAMP not null,
    updated_at       TIMESTAMP not null
);

-- +goose Down

DROP TABLE tokens;