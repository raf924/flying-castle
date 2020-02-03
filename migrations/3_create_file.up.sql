CREATE TABLE file
(
    id          integer primary key not null,
    name        text                not null,
    first_chunk_id integer             not null,
    path_id     integer             not null,
    size        integer             not null,
    modified_at timestamp           not null,
    foreign key (path_id) references path (id),
    foreign key (first_chunk_id) references chunk (id)
);