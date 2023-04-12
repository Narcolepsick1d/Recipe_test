

CREATE TABLE IF NOT EXISTS users(
                      id int primary key generated always as  identity ,
                      name varchar,
                      email varchar,
                      password varchar,
                      registered_at timestamp
);

CREATE TABLE IF NOT EXISTS refresh_tokens(
                              id serial not null unique ,
                              user_id int references users(id) on delete cascade not null ,
                              token varchar(255) not null unique ,
                              expires_at timestamp not null
);

CREATE TABLE IF NOT EXISTS recipe(
                       id int primary key generated always as  identity,
                       name varchar,
                       description text,
                       ingredients text,
                       total_time int
);
CREATE TABLE IF NOT EXISTS steps(
                      recipe_id int references recipe(id) on DELETE cascade ,
                      step_number int,
                      step_description text,
                      time_per_step int
);
CREATE TABLE IF NOT EXISTS rates(

                      recipe_id int references recipe(id) on DELETE cascade ,
                      rate int check ( rate>=1 and rate<=5),
                      rate_quantity int
);


INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Pasta with Tomato Sauce', 'Spaghetti with homemade tomato sauce', 'spaghetti, tomatoes, onion, garlic', 3000);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Tacos with Beef', 'Tacos with seasoned ground beef', 'ground beef, taco seasoning, tortillas, lettuce, cheese, sour cream', 6000);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Chicken Alfredo', 'Creamy Alfredo sauce with chicken and pasta', 'fettuccine pasta, chicken breasts, heavy cream, garlic, butter, parmesan cheese', 3500);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Roasted Vegetables Salad', 'Healthy salad with roasted veggies and balsamic dressing', 'mixed vegetables, balsamic vinegar, olive oil, salt, pepper', 2000);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Beef Stroganoff', 'Beef and mushroom sauce over egg noodles', 'beef sirloin, mushrooms, egg noodles, sour cream, onion, garlic', 9000);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Grilled Steak with Chimichurri Sauce', 'Juicy steak with zesty chimichurri sauce', 'rib-eye steak, parsley, cilantro, garlic, vinegar, olive oil', 30);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Spicy Shrimp Stir Fry', 'Stir fry with spicy shrimp and mixed vegetables', 'shrimp, mixed vegetables, soy sauce, garlic, ginger, sriracha', 25);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Baked Salmon', 'Baked salmon with lemon and herbs', 'salmon fillet, lemon, herbs, butter, salt, pepper', 20);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Beef Burritos', 'Burritos filled with seasoned ground beef', 'ground beef, taco seasoning, tortillas, cheddar cheese, salsa, lettuce, sour cream', 30);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Caprese Salad', 'Fresh salad with tomato, mozzarella, and basil', 'tomatoes, fresh mozzarella cheese, basil leaves, olive oil, balsamic glaze', 15);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Chicken Fajitas', 'Fajitas with marinated chicken and bell peppers', 'chicken breasts, bell peppers, onions, lime, garlic, cumin, chili powder', 25);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Pesto Pasta Salad', 'Pasta salad with pesto, sun-dried tomatoes, and feta', 'penne pasta, pesto sauce, sun-dried tomatoes, feta cheese, olive oil', 20);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Beef and Broccoli Stir Fry', 'Stir fry with beef and broccoli in a garlic sauce', 'beef sirloin, broccoli, soy sauce, garlic, ginger, cornstarch', 25);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Caesar Salad', 'Classic salad with romaine lettuce and Caesar dressing', 'romaine lettuce, croutons, parmesan cheese, Caesar dressing', 15);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Barbecue Chicken Pizza', 'Pizza with barbecue sauce and chicken', 'pizza crust, barbecue sauce, cooked chicken, red onion, mozzarella cheese', 30);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Thai Chicken Curry', 'Curry with chicken, vegetables, and coconut milk', 'chicken breasts, green curry paste, coconut milk, mixed vegetables, fish sauce', 35);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Stuffed Bell Peppers', 'Bell peppers stuffed with ground turkey and rice', 'bell peppers, ground turkey, rice, tomato sauce, onion, garlic', 40);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Baked Ziti', 'Pasta with tomato sauce and cheese', 'ziti pasta, tomato sauce, ricotta cheese, mozzarella cheese', 25);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Honey Mustard Glazed Salmon', 'Salmon with a sweet and tangy honey mustard glaze', 'salmon fillet, honey, dijon mustard, soy sauce, garlic', 20);

INSERT INTO recipe (name, description, ingredients, total_time) VALUES ('Vegetarian Chili', 'Chili with mixed vegetables and beans', 'mixed vegetables, beans, tomato sauce, chili powder, cumin, onion, garlic', 35);


INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (1,5,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (2,1,'take taco and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (2,2,'take beef ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (2,3,'take lettuce and ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (2,4,'take plate and ...',2000);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (2,4,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (3,1,'take chicken and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (3,2,'take pasta ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (3,3,'take plate and ...',1000);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (3,2,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (4,1,'take cucumber and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (4,2,'take tomato ...',1000);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (4,5,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (5,1,'take beef and ...',3000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (5,2,'take tomato ...',3500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (5,3,'take pan and fry meat ...',500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (5,4,'take sauce and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (5,5,'take plate and ...',1000);

INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (5,2,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (6,1,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (7,5,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (8,5,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (9,4,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (10,4,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (11,3,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (12,3,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (13,2,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (14,5,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (15,2,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (16,4,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (17,5,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (18,5,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (19,3,1);

INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,1,'take spaghetti and ...',1000);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,2,'take tomato ...',1500);
INSERT INTO steps(recipe_id,step_number,step_description,time_per_step) VALUES (1,3,'take plate and ...',500);
INSERT INTO rates(recipe_id,rate,rate_quantity)VALUES (20,5,1);