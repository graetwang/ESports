/*
Navicat MySQL Data Transfer

Source Server         : esports-service
Source Server Version : 50743
Source Host           : 139.9.113.33:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50743
File Encoding         : 65001

Date: 2024-07-16 09:25:40
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for schedules
-- ----------------------------
DROP TABLE IF EXISTS `schedules`;
CREATE TABLE `schedules` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `matchId` bigint(20) DEFAULT NULL,
  `matchTime` datetime(3) DEFAULT NULL,
  `round` varchar(191) DEFAULT NULL,
  `weight` bigint(20) DEFAULT NULL,
  `match_event` varchar(191) DEFAULT NULL,
  `project` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_schedules_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;