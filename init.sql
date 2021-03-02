CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS  candies (
    candy_id uuid DEFAULT uuid_generate_v4(),
    category VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    price NUMERIC NOT NULL,
    PRIMARY KEY (candy_id)
);