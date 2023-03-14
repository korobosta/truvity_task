CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    created timestamp
);