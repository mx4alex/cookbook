CREATE TABLE cousine(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT
);

CREATE TABLE category(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT
);

CREATE TABLE dish(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    category_id INTEGER REFERENCES category("id"),
    cousine_id INTEGER REFERENCES cousine("id"),
    description TEXT,
    recipe TEXT,
    time INTEGER
);

CREATE TABLE "ingredient"(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    measure_unit VARCHAR(255),
    protein INTEGER,
    fats INTEGER,
    carbohydrates INTEGER
);

CREATE TABLE dish_ingredient(
    id SERIAL PRIMARY KEY,
    dish_id INTEGER REFERENCES dish("id"),
    ingredient_id INTEGER REFERENCES ingredient("id"),
    quantity INTEGER
);
