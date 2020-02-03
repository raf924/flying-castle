CREATE TABLE storage_key
(
    id         integer primary key not null,
    key        TEXT                not null unique,
    created_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP
);