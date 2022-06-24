CREATE TABLE records
(
    id serial not null unique,
    long_url varchar(255) not null unique,
    token varchar(255) not null unique,
    created_at timestamp not null,
    expiry_date timestamp not null
);