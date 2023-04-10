CREATE table users(
                      id int primary key generated always as  identity ,
                      name varchar,
                      email varchar,
                      password varchar,
                      registered_at timestamp
);

CREATE table refresh_tokens(
                              id serial not null unique ,
                              user_id int references users(id) on delete cascade not null ,
                              token varchar(255) not null unique ,
                              expires_at timestamp not null
);