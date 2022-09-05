CREATE TABLE IF NOT EXISTS phrases (
    id SERIAL PRIMARY KEY,
    content VARCHAR NOT NULL,
    group_ VARCHAR (50) NOT NULL,
    author VARCHAR (50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);