
CREATE DATABASE HMS;

use HMS;

CREATE TABLE patient  (
     id         VARCHAR(50) PRIMARY KEY ,
     name       VARCHAR(50),
     age        INT,
     gender     VARCHAR(10),
     phone      VARCHAR(25),
     conditions VARCHAR(50)
  );
