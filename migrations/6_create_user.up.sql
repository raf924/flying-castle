CREATE TABLE user
(
    id                  integer primary key not null,
    user_name           text                not null unique,
    hash                text                not null,
    storage_key         text                not null,
    root_path_id        integer unique,
    main_group_id       integer unique,
    foreign key (root_path_id) references path  (id),
    foreign key (main_group_id) references 'group' (id)
);