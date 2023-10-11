CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) NOT NULL,
    nickname varchar(50) NOT NULL UNIQUE,
    email varchar(50) NOT NULL UNIQUE,
    password varchar(20) NOT NULL,
    createdAt timestamp DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;