--users table
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    token TEXT UNIQUE,        
    email VARCHAR(255) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,    
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('user', 'admin')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--article

CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--comment

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    article_id INT REFERENCES articles(id) ON DELETE CASCADE,
    user_id TEXT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
