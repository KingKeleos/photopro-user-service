create table if not exists users (
    index bigserial not null primary key,
    username varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    location_id bigserial,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)