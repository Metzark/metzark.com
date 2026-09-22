-- Verify metzark.com:tables/urls on pg

BEGIN;

SELECT 1
FROM information_schema.tables
WHERE table_schema = 'url_extender'
  AND table_name = 'urls';

ROLLBACK;
