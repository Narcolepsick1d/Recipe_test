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
CREATE table users(id int primary key generated always as  identity ,name varchar,email varchar,                      password varchar,registered_at timestamp);
CREATE table users(id int primary key generated always as  identity ,name varchar,email varchar,                      password varchar,registered_at timestamp);

CREATE table refresh_tokens(  id serial not null unique ,  user_id int references users(id) on delete cascade not null ,  token varchar(255) not null unique ,      expires_at timestamp not null);
CREATE table recipe(
    id int primary key generated always as  identity,
    name varchar,
    description text,
    ingredients text,
    steps text,
    total_time int,
    rates int,
    rates_quantity int
);

insert into recipe (name, description, ingredients, steps, total_time, rates, rates_quantity) values
('omlet','yes','milk,egg,butter','1.sdasdadasd. 2.sdasdadsa.',9000,5,1);


select * from recipe where ingredients like '%meat%';
drop table refresh_tokens;
drop table users;
drop table recipe;