CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

INSERT INTO companies (name) VALUES ('Naimix');