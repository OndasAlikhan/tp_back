create table races
(
    uuid        uuid primary key             default uuid_generate_v4(),
    status      varchar,
    lobby_id    uuid references lobbies (uuid),
    text        varchar,
    created_at timestamp           not null default current_timestamp,
    updated_at timestamp           not null default current_timestamp
);

create table race_users 
(
    user_id    uuid references users (uuid),
    race_id    uuid references races (uuid)
)
