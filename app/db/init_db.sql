CREATE TABLE viewing(
  id INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(250) NOT NULL,
  director VARCHAR(100) NOT NULL,
  year_made INT NOT NULL,
  date_watched DATE NOT NULL,
  PRIMARY KEY ( id )
);
