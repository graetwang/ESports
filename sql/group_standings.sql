/*
Navicat MySQL Data Transfer

Source Server         : esports-service
Source Server Version : 50743
Source Host           : 139.9.113.33:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50743
File Encoding         : 65001

Date: 2024-07-15 20:30:13
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for group_standings
-- ----------------------------
DROP TABLE IF EXISTS `group_standings`;
CREATE TABLE `group_standings` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` bigint(20) DEFAULT NULL,
  `team_id` bigint(20) DEFAULT NULL,
  `matches_played` bigint(20) DEFAULT NULL,
  `matches_won` bigint(20) DEFAULT NULL,
  `matches_lost` bigint(20) DEFAULT NULL,
  `points` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_group_standings_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
