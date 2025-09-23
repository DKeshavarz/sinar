CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone VARCHAR(20),
    profile_pic VARCHAR(255),
    student_num VARCHAR(50),
    sex BOOLEAN
);

CREATE TABLE universities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    location VARCHAR(255),
    logo VARCHAR(255)
);

CREATE TABLE restaurants (
    id SERIAL PRIMARY KEY,
    university_id INTEGER REFERENCES universities(id),
    name VARCHAR(255),
    sex BOOLEAN,
    color VARCHAR(50)
);

CREATE TABLE foods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE user_universities (
    user_id INTEGER REFERENCES users(id),
    university_id INTEGER REFERENCES universities(id),
    PRIMARY KEY (user_id, university_id)
);

CREATE TABLE user_foods (
    user_id INTEGER REFERENCES users(id),
    food_id INTEGER REFERENCES foods(id),
    restaurant_id INTEGER REFERENCES restaurants(id),
    price BIGINT,
    sinar_price BIGINT,
    code VARCHAR(20),
    ttl INTEGER,
    PRIMARY KEY (user_id, food_id, restaurant_id)
);