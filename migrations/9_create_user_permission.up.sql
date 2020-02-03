create table user_permission(
  path integer not null,
  user integer not null,
  read INTEGER(1) not null,
  write integer(1) not null,
  primary key (path, user),
  foreign key (path) references path(id),
  foreign key (user) references user(id)
);