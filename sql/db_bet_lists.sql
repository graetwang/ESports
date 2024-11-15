/*
Navicat MySQL Data Transfer

Source Server         : esports-service
Source Server Version : 50743
Source Host           : 139.9.113.33:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50743
File Encoding         : 65001

Date: 2024-07-29 09:31:44
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for db_bet_lists
-- ----------------------------
DROP TABLE IF EXISTS `db_bet_lists`;
CREATE TABLE `db_bet_lists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `list_id` bigint(20) unsigned DEFAULT NULL,
  `bet_id` longtext,
  `title` longtext,
  `category_name` longtext,
  `bet_end_time` longtext,
  `bet_end_time_txt` longtext,
  `status` longtext,
  `status_txt` longtext,
  `total_price` longtext,
  `people_num` longtext,
  `result_item_id` longtext,
  `dynamic_id` bigint(20) DEFAULT NULL,
  `team_a_win` longtext,
  `team_b_win` longtext,
  `match_status` longtext,
  `match_bet_count` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_db_bet_lists_deleted_at` (`deleted_at`),
  KEY `fk_db_lists_db_bet_lists` (`list_id`),
  CONSTRAINT `fk_db_lists_db_bet_lists` FOREIGN KEY (`list_id`) REFERENCES `db_lists` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;
