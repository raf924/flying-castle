CREATE TABLE folder (
    id      integer primary key     not null,
    path_id integer unique          not null,
    foreign key (path_id) references path(id)
);