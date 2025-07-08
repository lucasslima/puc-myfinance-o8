CREATE TABLE `entry` (
  `code` int(11) NOT NULL AUTO_INCREMENT,
  `description` varchar(100) DEFAULT NULL,
  `date` datetime NOT NULL,
  `value` float NOT NULL,
  PRIMARY KEY (`code`),
  CONSTRAINT `account` FOREIGN KEY (`code`) REFERENCES `accounts` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;