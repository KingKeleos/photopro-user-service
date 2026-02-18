create table if not exists permissions(
    index bigserial not null primary key,
    user_id bigserial not null,
    project_id bigserial not null,
    role_id bigserial not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)