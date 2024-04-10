create table tasks
(
    uuid        uuid primary key             default uuid_generate_v4(),
    project_id   uuid references projects (uuid),
    title       varchar(255) unique not null,
    description varchar(500)                 default null,
    priority    int,
    status      int,
    created_at timestamp           not null default current_timestamp,
    updated_at timestamp           not null default current_timestamp
)