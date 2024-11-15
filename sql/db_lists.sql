/*
Navicat MySQL Data Transfer

Source Server         : esports-service
Source Server Version : 50743
Source Host           : 139.9.113.33:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50743
File Encoding         : 65001

Date: 2024-07-29 09:31:31
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for db_lists
-- ----------------------------
DROP TABLE IF EXISTS `db_lists`;
CREATE TABLE `db_lists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `v` longtext,
  `key` longtext,
  `new_day` longtext,
  `date_txt` longtext,
  `view_type` bigint(20) DEFAULT NULL,
  `tournament_id` longtext,
  `tournament_name` longtext,
  `tournament_en_name` longtext,
  `tournament_image` longtext,
  `round_name` longtext,
  `round_son_name` longtext,
  `match_id` longtext,
  `team_image_thumb_a` longtext,
  `team_image_thumb_b` longtext,
  `match_date` longtext,
  `match_date1` longtext,
  `match_team_a` longtext,
  `match_team_b` longtext,
  `game_count` longtext,
  `start_id` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_db_lists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;
