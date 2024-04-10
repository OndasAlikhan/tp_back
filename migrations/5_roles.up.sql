create table roles
(
    uuid        uuid primary key default uuid_generate_v4(),
    name        varchar(255) unique not null,
    description varchar(500)     default null
)