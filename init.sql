CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS  candies (
    id uuid DEFAULT uuid_generate_v4(),
    category VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    price NUMERIC(6,4) NOT NULL,
    PRIMARY KEY (id)
);