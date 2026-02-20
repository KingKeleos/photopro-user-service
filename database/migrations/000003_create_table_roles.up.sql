create table if not exists roles(
    index bigserial not null primary key,
    name varchar(255) not null unique,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)