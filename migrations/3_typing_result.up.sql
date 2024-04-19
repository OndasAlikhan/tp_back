create table typing_result
(
    uuid        uuid primary key            default uuid_generate_v4(),
    duration        float                   default null,
    wpm             float                   default null,
    accuracy        float                   default null,
    user_id         uuid                    not null,
    race_id         uuid                    default null,
    created_at timestamp           not null default current_timestamp,
    updated_at timestamp           not null default current_timestamp
)