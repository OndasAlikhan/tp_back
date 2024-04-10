create table projects
(
    uuid        uuid primary key             default uuid_generate_v4(),
    name        varchar(255) unique not null,
    description varchar(500)                 default null,
    created_at timestamp           not null default current_timestamp,
    updated_at timestamp           not null default current_timestamp
)