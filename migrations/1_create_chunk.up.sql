CREATE TABLE chunk
(
    id         integer primary key not null,
    path       text                NOT NULL,
    next_chunk integer,
    FOREIGN KEY (next_chunk) REFERENCES chunk (id)
);