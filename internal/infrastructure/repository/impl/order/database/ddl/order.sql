CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    ticket_id VARCHAR(255) NOT NULL,
    user_id BIGINT NOT NULL,
    price VARCHAR(10),
    created_at TIMESTAMP,
    CONSTRAINT fk_ticket
        FOREIGN KEY(ticket_id) 
        REFERENCES tickets(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);