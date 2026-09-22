-- Deploy metzark.com:schema_url_extender to pg

BEGIN;

CREATE SCHEMA IF NOT EXISTS url_extender;

COMMIT;
