create table permissions
(
    uuid        uuid primary key default uuid_generate_v4(),
    role_id      uuid references roles (uuid),
    name        varchar(255) unique not null,
    description varchar(500)     default null
)