CREATE TABLE api_key
(
    id         integer primary key not null,
    key        text                not null,
    user       integer             not null,
    created_at timestamp           not null,
    foreign key (user) references user (id)
);