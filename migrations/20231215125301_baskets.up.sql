CREATE TABLE baskets
(
    id         SERIAL PRIMARY KEY,
    data       JSONB CHECK (octet_length(data::text) <= 2048),
    state      TEXT      NOT NULL,
    user_id    INTEGER REFERENCES users (id),
    created_at timestamp not null default now(),
    updated_at timestamp
);
