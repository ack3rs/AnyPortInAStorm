CREATE USER 'sqluser'@'%' IDENTIFIED BY 'sqlpass' ;
GRANT ALL ON *.* TO 'sqluser'@'%' WITH GRANT OPTION ;

CREATE DATABASE `ports`;

USE ports;

CREATE TABLE `portnames` (
  `portname` varchar(10) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `city` varchar(50) DEFAULT NULL,
  `country` varchar(50) DEFAULT NULL,
  `alias` varchar(50) DEFAULT NULL,
  `regions` varchar(50) DEFAULT NULL,
  `lat` double DEFAULT NULL,
  `lon` double DEFAULT NULL,
  `province` varchar(50) DEFAULT NULL,
  `timezone` varchar(50) DEFAULT NULL,
  `unlocs` varchar(50) DEFAULT NULL,
  `code` varchar(50) DEFAULT NULL,
  UNIQUE KEY `portname` (`portname`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

