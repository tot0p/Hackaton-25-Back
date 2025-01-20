

CREATE TABLE IF NOT EXISTS Users (
    UUID VARCHAR(36) PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS JWT (
    UUID VARCHAR(36) PRIMARY KEY,
    userUUID VARCHAR(36) NOT NULL,
    token VARCHAR(255) NOT NULL,
    ExpiresAt TIMESTAMP NOT NULL,
    FOREIGN KEY (userUUID) REFERENCES users(UUID)
);


CREATE TABLE IF NOT EXISTS Travels (
    UUID VARCHAR(36) PRIMARY KEY,
    userUUID VARCHAR(36) NOT NULL,
    Date DATE NOT NULL,
    StartLocation VARCHAR(255) NOT NULL,
    EndLocation VARCHAR(255) NOT NULL,
    Distance FLOAT NOT NULL,
    Duration FLOAT NOT NULL,
    CO2 FLOAT NOT NULL,
    FOREIGN KEY (userUUID) REFERENCES users(UUID)
);