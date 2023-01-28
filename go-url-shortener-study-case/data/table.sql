CREATE TABLE items (
    id BIGSERIAL PRIMARY KEY,
    long_url VARCHAR NOT NULL,
    slug VARCHAR NOT NULL UNIQUE,
    title VARCHAR,
    content VARCHAR,
    view_count BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE item_views (
    id BIGSERIAL PRIMARY KEY,
    item_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

ALTER TABLE public.items ADD CONSTRAINT items_slug_key UNIQUE (slug);

ALTER TABLE public.items DROP CONSTRAINT items_slug_key;