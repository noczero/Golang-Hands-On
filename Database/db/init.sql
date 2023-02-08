use zero_db;

show tables;

CREATE TABLE test (
                      id int NOT NULL,
                      dummy VARCHAR(100),
                      PRIMARY KEY (id)
) ENGINE=InnoDB;
show tables;
DESC test;

CREATE TABLE customer
(
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
) ENGINE = InnoDB;

DESC customer;

SELECT * FROM customer;

ALTER TABLE customer
    ADD email VARCHAR(100) NULL;

ALTER TABLE customer
    ADD balance DOUBLE NOT NULL;

ALTER TABLE customer
    ADD rating INT NOT NULL;

ALTER TABLE customer
    ADD birthDate TIMESTAMP NULL;

ALTER TABLE customer
DROP COLUMN birthDate;

ALTER TABLE customer
    ADD birth_date TIMESTAMP NULL;

ALTER TABLE customer
    ADD created_at TIMESTAMP DEFAULT current_timestamp;

ALTER TABLE customer
    ADD married BOOL NOT NULL;


INSERT INTO customer (id, name, email, balance, rating, birth_date, married)
VALUES ('budi', 'Budi', 'budi@test.com',1000.10,8,'2020-10-10',false),
       ('pratama', 'Pratama', 'pratama@test.com',2000.20,9,'2020-11-11',true),
       ('zero', 'Zero', 'zero@test.com',3000.30,10,'2020-9-9',false);

SELECT * FROM customer;

CREATE TABLE user (
                      id INT AUTO_INCREMENT,
                      username VARCHAR(100),
                      password VARCHAR(100),
                      PRIMARY KEY (id)
) ENGINE = InnoDB;
DESC user;

INSERT INTO user (username, password) VALUES ('admin','admin');
SELECT * FROM user;

# ADD comments table
CREATE TABLE comments(
                         id BIGINT AUTO_INCREMENT NOT NULL,
                         email VARCHAR(100),
                         comment TEXT,
                         PRIMARY KEY (id)
) ENGINE = InnoDB;
# DROP TABLE comments;