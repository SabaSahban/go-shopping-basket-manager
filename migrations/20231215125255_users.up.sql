CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   TEXT      NOT NULL,
    password   TEXT      NOT NULL,
    created_at timestamp not null default now(),
    updated_at timestamp
);
