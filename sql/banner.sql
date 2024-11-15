/*
Navicat MySQL Data Transfer

Source Server         : ESports
Source Server Version : 50744
Source Host           : 192.168.249.142:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50744
File Encoding         : 65001

Date: 2024-07-10 19:19:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `picture` varchar(191) DEFAULT NULL,
  `title` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_banner_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
