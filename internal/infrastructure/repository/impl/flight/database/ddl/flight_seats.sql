CREATE TABLE flight_seats (
    seat_id VARCHAR(10) PRIMARY KEY,
    flight_id VARCHAR(255) NOT NULL,
    class VARCHAR(50) NOT NULL,
    is_over_sold BOOLEAN NOT NULL DEFAULT false,
    available BOOLEAN NOT NULL DEFAULT true,
    CONSTRAINT fk_flight
        FOREIGN KEY (flight_id)
        REFERENCES flight (flight_id)
        ON DELETE CASCADE
);