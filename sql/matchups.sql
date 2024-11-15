/*
Navicat MySQL Data Transfer

Source Server         : ESports
Source Server Version : 50744
Source Host           : 192.168.249.142:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50744
File Encoding         : 65001

Date: 2024-07-10 15:02:41
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for matchups
-- ----------------------------
DROP TABLE IF EXISTS `matchups`;
CREATE TABLE `matchups` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `matchId` bigint(20) DEFAULT NULL,
  `homeId` bigint(20) DEFAULT NULL,
  `awayId` bigint(20) DEFAULT NULL,
  `startTime` datetime(3) DEFAULT NULL,
  `endTime` datetime(3) DEFAULT NULL,
  `winnerId` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_matchups_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
