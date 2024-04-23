CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT encode(sha256(random()::text::bytea), 'hex')
);

CREATE TABLE IF NOT EXISTS projects(
    id SERIAL PRIMARY KEY ,
    name TEXT UNIQUE NOT NULL,
    content JSON NOT NULL,
    description TEXT DEFAULT NULL
);