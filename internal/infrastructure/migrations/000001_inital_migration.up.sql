CREATE TABLE users (
    id       SERIAL PRIMARY KEY,
    username TEXT,
    email    TEXT
);


CREATE TABLE addresses (
    id          SERIAL PRIMARY KEY,
    user_id     BIGINT REFERENCES users (id),
    street      TEXT,
    city        TEXT,
    state       TEXT,
    country     TEXT,
    postal_code TEXT
);






