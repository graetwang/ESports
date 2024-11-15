/*
Navicat MySQL Data Transfer

Source Server         : ESports
Source Server Version : 50744
Source Host           : 192.168.249.142:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50744
File Encoding         : 65001

Date: 2024-07-10 15:02:52
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for players
-- ----------------------------
DROP TABLE IF EXISTS `players`;
CREATE TABLE `players` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `player_id` bigint(20) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `nickname` varchar(191) DEFAULT NULL,
  `role` varchar(191) DEFAULT NULL,
  `teamId` bigint(20) DEFAULT NULL,
  `photo` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_players_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
