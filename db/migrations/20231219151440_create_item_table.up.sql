-- Define the enum types
CREATE TYPE category_enum AS ENUM ('Weapon', 'Artifact', 'Potion');
CREATE TYPE rarity_enum AS ENUM ('Common', 'Uncommon', 'Rare', 'Epic', 'Legendary');

-- Create the items table
CREATE TABLE IF NOT EXISTS items (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       category category_enum NOT NULL,
                       rarity rarity_enum NOT NULL,
                       cost DOUBLE PRECISION NOT NULL,
                       description TEXT NOT NULL
);

-- Insert dummy data into the items table
INSERT INTO items (name, category, rarity, cost, description) VALUES
                                                                  ('Sword of Power', 'Weapon', 'Epic', 100.0, 'A mighty sword with magical powers.'),
                                                                  ('Healing Potion', 'Potion', 'Rare', 25.0, 'A potion that restores health.'),
                                                                  ('Ancient Scroll', 'Artifact', 'Legendary', 500.0, 'A mysterious scroll with ancient knowledge.'),
                                                                  ('Bow of Precision', 'Weapon', 'Uncommon', 75.0, 'A bow that enhances accuracy.'),
                                                                  ('Invisibility Cloak', 'Artifact', 'Legendary', 1000.0, 'A cloak that grants invisibility.');

-- Add more dummy data as needed
