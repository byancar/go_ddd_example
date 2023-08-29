CREATE TABLE user (
    id       SERIAL PRIMARY KEY,
    username TEXT,
    email    TEXT
);


CREATE TABLE address (
    id          SERIAL PRIMARY KEY,
    user_id     BIGINT REFERENCES user (id),
    street      TEXT,
    city        TEXT,
    state       TEXT,
    country     TEXT,
    postal_code TEXT
);






