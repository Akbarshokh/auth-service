BEGIN;

CREATE TABLE IF NOT EXISTS users(
    "id" serial PRIMARY KEY,
    "client_id" varchar(40) UNIQUE NOT NULL,
    "first_name" varchar(255) NOT NULL,
    "last_name" varchar(255) NOT NULL,
    "email" varchar(50) NOT NULL,
    "device_num" varchar(100) NOT NULL,
    "device_type" varchar(100) NOT NULL,
    "active" boolean,
    "access_token" varchar NOT NULL,
    "refresh_token" varchar NOT NULL
);

COMMIT;