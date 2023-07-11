-- +goose Up
CREATE ROLE decision LOGIN PASSWORD 'decision' NOINHERIT CREATEDB;
CREATE SCHEMA decision AUTHORIZATION decision;
GRANT USAGE ON SCHEMA decision TO PUBLIC;

-- +goose Down
DROP SCHEMA decision;
DROP ROLE decision;
