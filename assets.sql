-- MySQL dump 10.13  Distrib 8.0.15, for osx10.14 (x86_64)
--
-- Host: localhost    Database: assets
-- ------------------------------------------------------
-- Server version	8.0.15

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `assets`
--

DROP TABLE IF EXISTS `assets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `assets` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type_name` varchar(32) NOT NULL DEFAULT '0',
  `cost` decimal(19,2) NOT NULL DEFAULT '0.00',
  `market_value` decimal(19,2) NOT NULL DEFAULT '0.00',
  `earnings` decimal(19,2) NOT NULL DEFAULT '0.00',
  `roa` decimal(19,2) NOT NULL DEFAULT '0.00',
  `ratio` decimal(19,2) NOT NULL DEFAULT '0.00',
  `details` varchar(255) NOT NULL DEFAULT '',
  `remark` varchar(255) NOT NULL DEFAULT '',
  `ctime` int(11) NOT NULL DEFAULT '0',
  `mtime` int(11) NOT NULL DEFAULT '0',
  `uid` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `type_name_2` (`type_name`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets`
--

LOCK TABLES `assets` WRITE;
/*!40000 ALTER TABLE `assets` DISABLE KEYS */;
INSERT INTO `assets` VALUES (10,'cash',76627.00,76800.00,173.00,0.23,34.61,'招商银行活期，余额宝','',1578152366,1578152366,0),(16,'gold',32000.00,35200.00,3200.00,10.00,14.45,'积存金，博士黄金ETF','',1578130828,1578130828,0),(17,'btc',20000.00,8000.00,-12000.00,-60.00,9.03,'比特币','',1578152273,1578152273,0),(18,'stock',63000.00,53605.00,-9395.00,-14.91,28.46,'哔哩哔哩，小米集团','',1578152990,1578152990,0),(19,'fund',29774.00,30817.64,1044.00,3.51,13.45,'广发聚鑫债券，兴全社会责任混合','',1578152703,1578152703,0);
/*!40000 ALTER TABLE `assets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `assets_log`
--

DROP TABLE IF EXISTS `assets_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `assets_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type_name` varchar(32) NOT NULL DEFAULT '0',
  `cost` decimal(19,2) NOT NULL DEFAULT '0.00',
  `market_value` decimal(19,2) NOT NULL DEFAULT '0.00',
  `earnings` decimal(19,2) NOT NULL DEFAULT '0.00',
  `roa` decimal(19,2) NOT NULL DEFAULT '0.00',
  `ratio` decimal(19,2) NOT NULL DEFAULT '0.00',
  `details` varchar(255) NOT NULL DEFAULT '',
  `remark` varchar(255) NOT NULL DEFAULT '',
  `ctime` int(11) NOT NULL DEFAULT '0',
  `uid` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `type_name` (`type_name`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets_log`
--

LOCK TABLES `assets_log` WRITE;
/*!40000 ALTER TABLE `assets_log` DISABLE KEYS */;
INSERT INTO `assets_log` VALUES (10,'cash',76627.00,76800.00,173.00,0.23,34.61,'招商银行活期，余额宝','',1578152366,0),(16,'gold',32000.00,35200.00,3200.00,10.00,14.45,'积存金，博士黄金ETF','',1578130828,0),(17,'btc',20000.00,8000.00,-12000.00,-60.00,9.03,'比特币','',1578152273,0),(18,'stock',63000.00,53605.00,-9395.00,-14.91,28.46,'哔哩哔哩，小米集团','',1578152990,0),(19,'fund',29773.00,30817.64,1044.64,3.51,13.45,'广发聚鑫债券，兴全社会责任混合','',1578130873,0);
/*!40000 ALTER TABLE `assets_log` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-01-05 10:08:33
