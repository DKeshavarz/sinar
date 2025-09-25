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