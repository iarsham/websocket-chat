-- 000002_init_schema.up.sql
-- This file contains the SQL statements to create the rooms table and its columns

CREATE TABLE IF NOT EXISTS rooms (
    id        SERIAL       PRIMARY KEY,
    name      varchar(255) NOT NULL
);

CREATE INDEX ix_name ON rooms (name);

