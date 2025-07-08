DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `code` int(11) NOT NULL AUTO_INCREMENT,
  `description` varchar(100) NOT NULL,
  `date` datetime DEFAULT NULL,
  `account` int(11) DEFAULT NULL,
  `value` float DEFAULT NULL,
  PRIMARY KEY (`code`),
  KEY `transactions_accounts_FK` (`account`),
  CONSTRAINT `transactions_accounts_FK` FOREIGN KEY (`account`) REFERENCES `accounts` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;