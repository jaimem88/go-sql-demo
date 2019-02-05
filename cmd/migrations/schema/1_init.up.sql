CREATE SCHEMA demo;
CREATE extension IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS demo.user (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL DEFAULT '',
    mobile TEXT, -- nullable
    age BIGINT, -- nullable
    admin BOOLEAN NOT NULL DEFAULT false,
);


CREATE TABLE IF NOT EXISTS demo.address (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    street_address TEXT NOT NULL DEFAULT '',
    suburb TEXT NOT NULL DEFAULT '',
    postcode TEXT NOT NULL DEFAULT '',
    state TEXT NOT NULL DEFAULT '',
    country TEXT NOT NULL DEFAULT '',    
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES user(id)
);
