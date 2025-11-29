-- – PostgreSQL schema for users table with ENUM group_type

-- Create ENUM
CREATE TYPE group_type AS ENUM
    ( 'IT', 'HR', 'PR', 'IO', 'JFR', 'Grafika' );

-- Create Table
CREATE TABLE users ( id SERIAL PRIMARY KEY,
first_name TEXT NOT NULL, last_name TEXT NOT NULL, email TEXT NOT NULL,
“group” group_type NOT NULL );

-- Example Insert
INSERT INTO users
    (first_name, last_name, email, “group”)
VALUES
    ('Jan', 'Kowalski', 'jan.kowalski@iaeste.pl', 'IT'),
    ('Alicja', 'Nowak', 'alicja.nowak@iaeste.pl', 'Grafika');
