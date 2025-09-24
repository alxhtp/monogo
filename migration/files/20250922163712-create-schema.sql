
-- +migrate Up
CREATE SCHEMA IF NOT EXISTS "monogo";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate Down
DROP SCHEMA IF EXISTS "monogo";
DROP EXTENSION IF EXISTS "uuid-ossp";