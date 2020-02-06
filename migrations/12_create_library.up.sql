create table library (
    id integer primary key not null,
    folder_id integer unique not null,
    description text
);