CREATE TABLE IF NOT EXISTS github_insights (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    repositories INT NOT NULL,
    total_stars INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);