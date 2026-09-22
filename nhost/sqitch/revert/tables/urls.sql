-- Revert metzark.com:tables/urls from pg

BEGIN;

DROP TABLE IF EXISTS url_extender.urls;

COMMIT;
