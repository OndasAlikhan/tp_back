create table lobbies
(
    uuid          uuid primary key default uuid_generate_v4(),
    admin_user_id uuid references users (uuid),
    status        varchar(255)     not null,
    name          varchar(255)     not null
);

create table lobby_users
(
    user_id  uuid references users (uuid),
    lobby_id uuid references lobbies (uuid)
)