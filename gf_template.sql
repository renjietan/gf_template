/*
 Navicat Premium Data Transfer

 Source Server         : gf
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : localhost:3306
 Source Schema         : gf_template

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 30/12/2025 17:52:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for f_famliy
-- ----------------------------
DROP TABLE IF EXISTS `f_famliy`;
CREATE TABLE `f_famliy`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `delete_at` tinyint NOT NULL DEFAULT 0,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id` DESC) USING BTREE,
  CONSTRAINT `f_famliy_chk_1` CHECK (`delete_at` in (0,1))
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of f_famliy
-- ----------------------------
INSERT INTO `f_famliy` VALUES (2, 'f_SSasaki Yamato', 0, '2025-12-29 09:05:56', '2025-12-29 09:05:56');
INSERT INTO `f_famliy` VALUES (1, 'f_Virginia Cole', 0, '2025-12-29 09:06:02', '2025-12-29 09:06:02');

-- ----------------------------
-- Table structure for t_info
-- ----------------------------
DROP TABLE IF EXISTS `t_info`;
CREATE TABLE `t_info`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `Info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `delete_at` tinyint NOT NULL DEFAULT 0,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id` DESC) USING BTREE,
  CONSTRAINT `delete` CHECK (`delete_at` in (0,1))
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_info
-- ----------------------------
INSERT INTO `t_info` VALUES (35, '张三--详情', 0, '2025-12-30 14:56:50', '2025-12-30 14:56:50');
INSERT INTO `t_info` VALUES (5, 'i_OuNygb2PJ2', 0, '2025-12-29 09:05:26', '2025-12-29 09:05:26');
INSERT INTO `t_info` VALUES (4, 'i_jL7swgO3jA', 0, '2025-12-29 09:05:30', '2025-12-29 09:05:30');
INSERT INTO `t_info` VALUES (3, 'i_ssLub1EQZF', 0, '2025-12-29 09:05:33', '2025-12-29 09:05:33');
INSERT INTO `t_info` VALUES (2, 'i_5v8JfupEyt', 0, '2025-12-29 09:05:36', '2025-12-29 09:05:36');
INSERT INTO `t_info` VALUES (1, 'i_TbA0Lu5UTj', 0, '2025-12-29 09:05:39', '2025-12-29 09:05:39');

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `f_id` int NOT NULL,
  `infoId` int NOT NULL,
  `delete_at` tinyint NOT NULL DEFAULT 0,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id` DESC) USING BTREE,
  INDEX `user_to_famly`(`f_id` ASC) USING BTREE,
  INDEX `user_to_info`(`infoId` ASC) USING BTREE,
  CONSTRAINT `user_to_famly` FOREIGN KEY (`f_id`) REFERENCES `f_famliy` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `user_to_info` FOREIGN KEY (`infoId`) REFERENCES `t_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `t_user_chk_1` CHECK (`delete_at` in (0,1))
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES (30, '张三', 1, 35, 0, '2025-12-30 14:56:50', '2025-12-30 14:56:50');
INSERT INTO `t_user` VALUES (5, 'u_Mao Yuning', 2, 5, 0, '2025-12-29 09:04:42', '2025-12-29 09:04:42');
INSERT INTO `t_user` VALUES (4, 'u_Brenda Nichols', 1, 4, 0, '2025-12-29 09:04:49', '2025-12-29 09:04:49');
INSERT INTO `t_user` VALUES (3, 'u_Yoshida Sakura', 1, 3, 0, '2025-12-29 09:04:54', '2025-12-29 09:04:54');
INSERT INTO `t_user` VALUES (2, 'u_Takada Kaito', 2, 2, 0, '2025-12-29 09:05:00', '2025-12-29 09:05:00');
INSERT INTO `t_user` VALUES (1, 'u_Cai Zhiyuan', 1, 1, 0, '2025-12-29 09:05:09', '2025-12-29 09:05:09');

SET FOREIGN_KEY_CHECKS = 1;
