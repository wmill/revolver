-- Deploy revolver:appschema to pg

BEGIN;

CREATE SCHEMA IF NOT EXISTS revolver_user;

COMMIT;
