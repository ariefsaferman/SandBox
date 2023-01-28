CREATE TABLE articles (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR,
    content VARCHAR,
    status VARCHAR(16),
    category_id INTEGER
);

CREATE TABLE tags (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR,
    description VARCHAR
);

CREATE TABLE article_tags (
    id BIGSERIAL PRIMARY KEY,
    article_id BIGINT,
    tag_id BIGINT
);

CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR,
    description VARCHAR
);