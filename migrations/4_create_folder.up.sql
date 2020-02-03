CREATE TABLE folder (
    id      integer primary key not null,
    path_id integer             not null,
    name    text                not null,
    foreign key (path_id) references path(id)
);