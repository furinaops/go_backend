-- Synthetic seed data for the User struct (GORM table "users").
-- Run with: psql "postgres://username:password@localhost:5432/database?sslmode=disable" -f seed.sql

CREATE TABLE IF NOT EXISTS users (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    role TEXT NOT NULL
);

INSERT INTO users (id, name, role) VALUES
    (1,  'Aarav Sharma',    'admin'),
    (2,  'Priya Patel',     'developer'),
    (3,  'Rohan Mehta',     'designer'),
    (4,  'Ananya Iyer',     'manager'),
    (5,  'Vikram Singh',    'analyst'),
    (6,  'Neha Gupta',      'devops'),
    (7,  'Karan Malhotra',  'qa'),
    (8,  'Sneha Reddy',     'support'),
    (9,  'Arjun Nair',      'sales'),
    (10, 'Ishita Bose',     'intern')
ON CONFLICT (id) DO NOTHING;

-- Keep the BIGSERIAL sequence in sync after inserting explicit IDs.
SELECT setval(pg_get_serial_sequence('users', 'id'), (SELECT MAX(id) FROM users));
