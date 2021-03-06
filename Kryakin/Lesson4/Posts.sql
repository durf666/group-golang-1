-- Adminer 4.7.7 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP DATABASE IF EXISTS `PostList`;
CREATE DATABASE `PostList` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `PostList`;

DROP TABLE IF EXISTS `Posts`;
CREATE TABLE `Posts` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Header` varchar(100) NOT NULL,
  `Text` varchar(255) NOT NULL,
  `Date` datetime NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `Posts` (`ID`, `Header`, `Text`, `Date`) VALUES
(1,	'Post1',	'This is post number one! 123',	'2020-08-25 01:15:12'),
(2,	'Interesting post',	'This is another interesting post',	'2020-08-22 11:32:00'),
(3,	'Post3',	'This is a new poster',	'2020-08-24 22:22:26');

-- 2020-08-25 01:18:47
