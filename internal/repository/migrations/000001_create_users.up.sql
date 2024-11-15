CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    third_name VARCHAR,
    email VARCHAR NOT NULL UNIQUE,
    phone_number VARCHAR,
    password VARCHAR NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    role VARCHAR NOT NULL
);