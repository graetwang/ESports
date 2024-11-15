/*
Navicat MySQL Data Transfer

Source Server         : esports-service
Source Server Version : 50743
Source Host           : 139.9.113.33:3306
Source Database       : esports

Target Server Type    : MYSQL
Target Server Version : 50743
File Encoding         : 65001

Date: 2024-07-12 16:37:12
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `section` varchar(191) DEFAULT NULL,
  `title` varchar(191) DEFAULT NULL,
  `cover` varchar(191) DEFAULT NULL,
  `duration` varchar(191) DEFAULT NULL,
  `view_counts` bigint(20) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `author` varchar(191) DEFAULT NULL,
  `address` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_videos_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
