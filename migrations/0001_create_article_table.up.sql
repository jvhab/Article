CREATE TABLE IF NOT EXISTS article (
    title VARCHAR(50) NOT NULL PRIMARY KEY,
    description TEXT,
    body TEXT,
    counts int,
    tag_list TEXT
);