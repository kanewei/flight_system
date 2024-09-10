CREATE TABLE airplanes (
    id SERIAL PRIMARY KEY,
    model VARCHAR(255) NOT NULL,
    seats JSON NOT NULL,
    available BOOLEAN NOT NULL
);