CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    birthdate DATE,
    password_hash VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

-- Add an index on email and username for faster lookups
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
