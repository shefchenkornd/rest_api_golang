CREATE TABLE IF NOT EXISTS articles
(
    id      bigserial not null primary key,
    title   varchar   not null unique,
    author  varchar   not null,
    content varchar   not null
);