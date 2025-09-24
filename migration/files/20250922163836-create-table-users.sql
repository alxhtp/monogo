
-- +migrate Up
CREATE TABLE IF NOT EXISTS "monogo"."users" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "status" smallint NOT NULL DEFAULT 0,
    "metadata" JSONB NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz NULL
);

-- +migrate Down
DROP TABLE IF EXISTS "monogo"."users";