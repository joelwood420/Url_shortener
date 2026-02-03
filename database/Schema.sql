create table urls (
    id integer primary key,
    long_url text not null,
    tag text not null,
    short_url text not null,
    user_id integer references users(id)
);

create table users (
    id integer primary key,
    username text not null unique,
    password_hash text not null
);