CREATE DATABASE amazon_ec;
\
c amazon_ec;

CREATE TABLE IF NOT EXISTS "users"
(
    user_id      VARCHAR(36) PRIMARY KEY,
    user_name    VARCHAR(100) NOT NULL,
    phone_number VARCHAR(13)  NOT NULL,
    gender       VARCHAR(10)  NOT NULL,
    created_at   timestamp    NOT NULL,
    updated_at   timestamp    NOT NULL,
    deleted_at   timestamp    NOT NULL
);
