CREATE TABLE tickets (
    id VARCHAR(255) PRIMARY KEY,
    user_id BIGINT NOT NULL,
    flight_id VARCHAR(255) NOT NULL,
    seat_id VARCHAR(10) NOT NULL,
    price VARCHAR(10) NOT NULL,
    status INTEGER NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id) 
        REFERENCES users (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_flight
        FOREIGN KEY (flight_id) 
        REFERENCES flights (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_flight_seat
        FOREIGN KEY (seat_id) 
        REFERENCES flight_seats (id)
        ON DELETE CASCADE
);