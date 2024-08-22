-- +goose Up

    create table users(
        id UUID primary key,
        name TEXT not null,
        created_at TIMESTAMP not null,
        updated_at TIMESTAMP not null
    );

-- +goose Down

DROP TABLE users;