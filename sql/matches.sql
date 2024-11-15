/*
Navicat MySQL Data Transfer

Source Server         : esports-service
Source Server Version : 50743
Source Host           : 139.9.113.33:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50743
File Encoding         : 65001

Date: 2024-07-16 09:25:22
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for matches
-- ----------------------------
DROP TABLE IF EXISTS `matches`;
CREATE TABLE `matches` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `project_id` bigint(20) DEFAULT NULL,
  `start_time` datetime(3) DEFAULT NULL,
  `end_time` datetime(3) DEFAULT NULL,
  `status` varchar(191) DEFAULT NULL,
  `bo` bigint(20) DEFAULT NULL,
  `home_id` bigint(20) DEFAULT NULL,
  `away_id` bigint(20) DEFAULT NULL,
  `home_score` bigint(20) DEFAULT NULL,
  `away_score` bigint(20) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `group_id` bigint(20) DEFAULT NULL,
  `phase` varchar(191) DEFAULT NULL,
  `home_name` varchar(191) DEFAULT NULL,
  `away_name` varchar(191) DEFAULT NULL,
  `group` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_matches_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
