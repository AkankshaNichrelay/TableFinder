Create database tablefinder;

use tablefinder;

DROP TABLE IF EXISTS restaurants;
CREATE TABLE restaurants (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(100) NOT NULL,
  location     VARCHAR(255) NOT NULL,
  ratings      DECIMAL(3,2),
  cuisine     VARCHAR(50),
  PRIMARY KEY (`id`)
);

INSERT INTO restaurants
  (name, location, ratings, cuisine)
VALUES
  ('Anatolia', '123 Magazine St, New Orleans', 4.5, "Mediterranean"),
  ('Mantra', '335 Harvey Blvb, New Orleans', 4.8, "Indian");