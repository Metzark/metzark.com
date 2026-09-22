-- Deploy metzark.com:tables/urls to pg

BEGIN;

CREATE TABLE IF NOT EXISTS url_extender.urls(
    extended TEXT PRIMARY KEY,
    actual TEXT,
    require_check BOOLEAN DEFAULT false,
    expires_at TIMESTAMP DEFAULT now() + interval '24 hours',
    created_at TIMESTAMP DEFAULT now()
);

COMMIT;
