-- Verify metzark.com:schema_url_extender on pg

BEGIN;

SELECT 1
FROM pg_namespace
WHERE nspname = 'url_extender';

ROLLBACK;
