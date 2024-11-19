CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    activity_no VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status INT,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
