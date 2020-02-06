CREATE TABLE user
(
    id                  integer primary key not null,
    user_name           text                not null unique,
    hash                text                not null unique,
    salt                text                not null unique,
    root_folder_id      integer unique not null,
    main_group_id       integer unique not null,
    foreign key (root_folder_id) references folder  (id),
    foreign key (main_group_id) references "group" (id)
);