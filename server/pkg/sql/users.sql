-- 000001_init_schema.up.sql
-- This file contains the SQL statements to create the users table and its columns

CREATE TABLE IF NOT EXISTS users
(
    id        uuid PRIMARY KEY         DEFAULT gen_random_uuid(),
    username  varchar(255)             NOT NULL UNIQUE,
    password  varchar(255)             NOT NULL,
    joined_at timestamp with time zone NOT NULL DEFAULT now(),
    last_seen timestamp with time zone NOT NULL DEFAULT now(),
    verified  boolean                  NOT NULL DEFAULT FALSE
);

CREATE INDEX ix_username ON users (username);

