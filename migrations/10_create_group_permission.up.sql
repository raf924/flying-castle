create table group_permission (
  path integer not null,
  'group' integer not null,
  read INTEGER(1) not null,
  write integer(1) not null,
  primary key (path, 'group'),
  foreign key (path) references path (id),
  foreign key ('group') references 'group' (id)
);