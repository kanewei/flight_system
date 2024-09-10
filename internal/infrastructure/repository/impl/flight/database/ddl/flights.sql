CREATE TABLE flights (
    id             VARCHAR(255) PRIMARY KEY,
    departure      VARCHAR(10) NOT NULL,
    arrival        VARCHAR(10) NOT NULL,
    departure_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    arrival_time   TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    price          VARCHAR(10) NOT NULL,
    airport_code   VARCHAR(10) REFERENCES airports(code),
    airplane_id    INTEGER REFERENCES airplanes(id)
    CONSTRAINT check_times CHECK (departure_time < arrival_time)
);