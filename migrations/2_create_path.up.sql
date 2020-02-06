create table path
(
    id          integer primary key not null,
    parent_id   integer,
    name        text,
    created_at  timestamp not null,
    accessed_at timestamp not null,
    modified_at timestamp not null,
    unique (parent_id, name),
    foreign key (parent_id) references path (id)
);