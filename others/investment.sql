/*
 Navicat Premium Data Transfer

 Source Server         : mysql_local
 Source Server Type    : MySQL
 Source Server Version : 80035 (8.0.35)
 Source Host           : localhost:3306
 Source Schema         : investment

 Target Server Type    : MySQL
 Target Server Version : 80035 (8.0.35)
 File Encoding         : 65001

 Date: 26/01/2024 10:33:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for asset
-- ----------------------------
DROP TABLE IF EXISTS `asset`;
CREATE TABLE `asset`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `stock` float NULL DEFAULT NULL,
  `option` float NULL DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of asset
-- ----------------------------
INSERT INTO `asset` VALUES (1, 50, 50, '2023-12-01 14:26:54');
INSERT INTO `asset` VALUES (2, 55.2, 49.8, '2023-12-08 14:28:20');
INSERT INTO `asset` VALUES (3, 40.5, 35.6, '2023-12-15 14:29:40');
INSERT INTO `asset` VALUES (4, 55, 30, '2023-12-22 14:31:13');
INSERT INTO `asset` VALUES (5, 75.2, 66.6, '2023-12-29 14:31:50');
INSERT INTO `asset` VALUES (6, 88.8, 99.9, '2024-01-05 14:32:30');

SET FOREIGN_KEY_CHECKS = 1;
