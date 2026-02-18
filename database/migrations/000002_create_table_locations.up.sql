create table if not exists locations(
    index bigserial not null primary key,
    name varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    street varchar(255),
    street_number int,
    country varchar(255),
    federal_state varchar(255)
)