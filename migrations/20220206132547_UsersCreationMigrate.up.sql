CREATE TABLE IF NOT EXISTS users
(
    id       bigserial not null primary key,
    login    varchar   not null unique,
    password varchar   not null
);
