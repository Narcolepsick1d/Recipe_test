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

CREATE table recipe(
                       id int primary key generated always as  identity,
                       name varchar,
                       description text,
                       ingredients text,
                       total_time int
);
create table steps(
                      recipe_id int references recipe(id) on DELETE cascade ,
                      step_number int,
                      step_description text,
                      time_per_step int
);
create table rates(

                      recipe_id int references recipe(id) on DELETE cascade ,
                      rate int check ( rate>=1 and rate<=5),
                      rate_quantity int
);


insert into recipe (name, description, ingredients, steps, total_time, rates, rates_quantity) values
('omlet','yes','milk,egg,butter','1.sdasdadasd. 2.sdasdadsa.',9000,5,1);


drop table steps;
drop table recipe;
drop table rates;
