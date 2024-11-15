/*
Navicat MySQL Data Transfer

Source Server         : esports-service
Source Server Version : 50743
Source Host           : 139.9.113.33:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50743
File Encoding         : 65001

Date: 2024-07-29 09:32:01
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for db_items
-- ----------------------------
DROP TABLE IF EXISTS `db_items`;
CREATE TABLE `db_items` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `bet_item_id` longtext,
  `init_price` longtext,
  `price` longtext,
  `member_max_bet` longtext,
  `item_name` longtext,
  `item_name_en` longtext,
  `item_name_tw` longtext,
  `win_rate` longtext,
  `odds` longtext,
  `bet_list_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_db_items_deleted_at` (`deleted_at`),
  KEY `fk_db_bet_lists_db_items` (`bet_list_id`),
  CONSTRAINT `fk_db_bet_lists_db_items` FOREIGN KEY (`bet_list_id`) REFERENCES `db_bet_lists` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4;
