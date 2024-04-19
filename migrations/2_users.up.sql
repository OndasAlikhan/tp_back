create table users
(
    uuid        uuid primary key             default uuid_generate_v4(),
    email       varchar(255) unique          not null,
    name        varchar(255)                 not null,
    password    varchar(255)                 not null,
    created_at  timestamp                    not null default current_timestamp,
    updated_at  timestamp                    not null default current_timestamp
)