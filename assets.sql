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
INSERT INTO `assets` VALUES (10,'cash',102600.00,103188.00,588.00,0.57,37.09,'招商银行活期，余额宝','',1581908606,1581908606,0),(16,'gold',34000.00,36112.00,2112.00,6.21,12.29,'积存金，博士黄金ETF','',1581992798,1581992798,0),(17,'btc',20000.00,20001.00,1.00,0.01,7.23,'比特币','',1581664153,1581664153,0),(18,'stock',80000.00,90300.00,10300.00,12.88,28.92,'港股：中芯国际，腾讯股份；\n美股：吉利德科技，纳斯达克ETF，百度','',1581908176,1581908176,0),(19,'fund',40000.00,40330.00,330.00,0.83,14.46,'广发聚鑫债券，纳斯达克指数','',1581992870,1581992870,0);
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

--
-- Table structure for table `assets_type`
--

DROP TABLE IF EXISTS `assets_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `assets_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `uid` int(11) NOT NULL DEFAULT '0',
  `ctime` int(11) NOT NULL DEFAULT '0',
  `mtime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets_type`
--

LOCK TABLES `assets_type` WRITE;
/*!40000 ALTER TABLE `assets_type` DISABLE KEYS */;
/*!40000 ALTER TABLE `assets_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `work_day`
--

DROP TABLE IF EXISTS `work_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `work_day` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` int(10) unsigned NOT NULL,
  `project_id` int(10) unsigned NOT NULL,
  `amount1` int(11) NOT NULL DEFAULT '0',
  `amount2` int(11) NOT NULL DEFAULT '0',
  `amount3` int(11) NOT NULL DEFAULT '0',
  `amount4` int(11) NOT NULL DEFAULT '0',
  `amount5` int(11) NOT NULL DEFAULT '0',
  `amount6` int(11) NOT NULL DEFAULT '0',
  `amount7` int(11) NOT NULL DEFAULT '0',
  `ctime` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `group_id` (`group_id`),
  KEY `project_id` (`project_id`),
  KEY `ctime` (`ctime`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `work_day`
--

LOCK TABLES `work_day` WRITE;
/*!40000 ALTER TABLE `work_day` DISABLE KEYS */;
INSERT INTO `work_day` VALUES (1,1,1,10,10,10,10,10,10,10,1581943862),(2,2,2,130,230,340,340,340,340,40,1581943926),(3,1,2,0,2,1,10,0,0,0,1581946583),(4,3,1,0,12,3,16,0,0,0,1581946646);
/*!40000 ALTER TABLE `work_day` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `work_group`
--

DROP TABLE IF EXISTS `work_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `work_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `work_group`
--

LOCK TABLES `work_group` WRITE;
/*!40000 ALTER TABLE `work_group` DISABLE KEYS */;
INSERT INTO `work_group` VALUES (1,'A'),(2,'B'),(3,'C');
/*!40000 ALTER TABLE `work_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `work_project`
--

DROP TABLE IF EXISTS `work_project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `work_project` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `work_project`
--

LOCK TABLES `work_project` WRITE;
/*!40000 ALTER TABLE `work_project` DISABLE KEYS */;
INSERT INTO `work_project` VALUES (1,'crush'),(2,'mare');
/*!40000 ALTER TABLE `work_project` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-02-18 15:35:43
