-- Revert metzark.com:schema_url_extender from pg

BEGIN;

DROP SCHEMA IF EXISTS url_extender;

COMMIT;
