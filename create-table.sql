/*This command will create a new table and delete the old one if it exists */

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id         INT AUTO_INCREMENT NOT NULL,
  username   VARCHAR(255) NOT NULL,
  pass       VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);
