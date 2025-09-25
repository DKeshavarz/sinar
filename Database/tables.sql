CREATE TABLE universities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    location VARCHAR(255),
    logo VARCHAR(255)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone VARCHAR(20),
    profile_pic VARCHAR(255),
    student_num VARCHAR(50),
    sex BOOLEAN,
    university_id INTEGER REFERENCES universities(id)
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

CREATE TABLE user_foods (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    food_id INTEGER REFERENCES foods(id),
    restaurant_id INTEGER REFERENCES restaurants(id),
    price BIGINT,
    sinar_price BIGINT,
    code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
);