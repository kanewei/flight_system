CREATE TABLE airports (
    airport_code VARCHAR(10) PRIMARY KEY,
    country VARCHAR(100) NOT NULL,
    airport_name VARCHAR(150) NOT NULL,
    terminal VARCHAR(50)
    available BOOLEAN NOT NULL
);