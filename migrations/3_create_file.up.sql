CREATE TABLE file
(
    id          integer primary key not null,
    first_chunk_id integer unique   not null,
    path_id     integer unique      not null,
    size        integer             not null,
    modified_at timestamp           not null,
    foreign key (path_id) references path (id),
    foreign key (first_chunk_id) references chunk (id)
);

create unique index unique_path on file(path_id);