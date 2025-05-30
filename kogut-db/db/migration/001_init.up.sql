CREATE TABLE kogut (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INTEGER CHECK (age >= 0 AND age <= 30),
    sex BOOLEAN NOT NULL
);
