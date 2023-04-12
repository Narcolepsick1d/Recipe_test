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
                       total_time int,
                       rates int,
                       rates_quantity int
);
create table steps(
    recipe_id int references recipe(id) on DELETE cascade ,
    step_number int,
    step_description text,
    time_per_step int
);

insert into recipe (name, description, ingredients, steps, total_time, rates, rates_quantity) values
('omlet','yes','milk,egg,butter','1.sdasdadasd. 2.sdasdadsa.',9000,5,1);


select * from recipe where ingredients like '%meat%';
drop table refresh_tokens;
drop table users;
drop table recipe;

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
                        user_id int references users(id) unique ,
                      recipe_id int references recipe(id) on DELETE cascade ,
                      rate int check ( rate>=0 and rate<=5),
                      rate_quantity int
);
drop table steps;
drop table recipe;
drop table rates;
insert into recipe (name, description, ingredients, total_time) VALUES ('omlet','egg with milk','egg,butter,milk',6000);
insert into steps (recipe_id, step_number, step_description, time_per_step) VALUES (1,1,'bit 2 eggs and mix with milk',6000);
insert into rates (user_id,recipe_id, rate, rate_quantity) VALUES (1,1,5,1);
select recipe.name,recipe.description,recipe.ingredients,recipe.total_time,rates.rate,rates.rate_quantity from recipe join rates  on recipe.id = rates.recipe_id  WHERE recipe.ingredients like '%egg%' ;
