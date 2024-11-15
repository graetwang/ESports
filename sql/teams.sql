/*
Navicat MySQL Data Transfer

Source Server         : esports-service
Source Server Version : 50743
Source Host           : 139.9.113.33:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50743
File Encoding         : 65001

Date: 2024-07-15 20:18:52
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for teams
-- ----------------------------
DROP TABLE IF EXISTS `teams`;
CREATE TABLE `teams` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `foundedDate` datetime(3) DEFAULT NULL,
  `region` varchar(50) DEFAULT NULL,
  `logo` varchar(191) DEFAULT NULL,
  `shortName` varchar(191) DEFAULT NULL,
  `created_by` bigint(20) unsigned DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) unsigned DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint(20) unsigned DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`),
  KEY `idx_teams_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
