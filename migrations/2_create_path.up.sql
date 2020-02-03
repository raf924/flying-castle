create table path
(
    id          integer primary key not null,
    parent_id   integer,
    created_at  timestamp not null,
    accessed_at timestamp not null,
    modified_at timestamp not null,
    foreign key (parent_id) references path (id)
);