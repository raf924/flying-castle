CREATE TABLE user_group (
    user_id integer not null,
    group_id integer not null,
    primary key (user_id, group_id),
    foreign key (user_id) references user  (id),
    foreign key (group_id) references 'group' (id)
)