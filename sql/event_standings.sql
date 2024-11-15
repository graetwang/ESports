/*
Navicat MySQL Data Transfer

Source Server         : esports-service
Source Server Version : 50743
Source Host           : 139.9.113.33:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50743
File Encoding         : 65001

Date: 2024-07-17 14:33:12
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for event_standings
-- ----------------------------
DROP TABLE IF EXISTS `event_standings`;
CREATE TABLE `event_standings` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `event` varchar(191) DEFAULT NULL,
  `team_id` bigint(20) DEFAULT NULL,
  `team` varchar(191) DEFAULT NULL,
  `team_logo` varchar(191) DEFAULT NULL,
  `matches_won` bigint(20) DEFAULT NULL,
  `matches_lost` bigint(20) DEFAULT NULL,
  `damage_per_minute` double DEFAULT NULL,
  `gold_per_minute` double DEFAULT NULL,
  `creep_score_per_minute` double DEFAULT NULL,
  `small_dragon_control_rate` double DEFAULT NULL,
  `average_kills_per_game` double DEFAULT NULL,
  `average_assists_per_game` double DEFAULT NULL,
  `average_deaths_per_game` double DEFAULT NULL,
  `baron_control_rate` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_event_standings_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
