/*
 Navicat Premium Dump SQL

 Source Server         : GF 数据库
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : localhost:3306
 Source Schema         : gf_template

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 29/01/2026 01:31:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gf_famliy
-- ----------------------------
DROP TABLE IF EXISTS `gf_famliy`;
CREATE TABLE `gf_famliy`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id` DESC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gf_famliy
-- ----------------------------
INSERT INTO `gf_famliy` VALUES (1, 'fm_1', NULL, NULL);

-- ----------------------------
-- Table structure for gf_info
-- ----------------------------
DROP TABLE IF EXISTS `gf_info`;
CREATE TABLE `gf_info`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `delete_at` int NOT NULL DEFAULT 0,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id` DESC) USING BTREE,
  CONSTRAINT `gf_info_chk_1` CHECK (`delete_at` in (0,1))
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gf_info
-- ----------------------------

-- ----------------------------
-- Table structure for gf_sys_cron
-- ----------------------------
DROP TABLE IF EXISTS `gf_sys_cron`;
CREATE TABLE `gf_sys_cron`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `group_id` bigint NOT NULL COMMENT '分组ID',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '任务标题',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '任务方法',
  `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '函数参数',
  `pattern` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '表达式',
  `policy` bigint NOT NULL DEFAULT 1 COMMENT '策略',
  `count` bigint NOT NULL DEFAULT 0 COMMENT '执行次数',
  `sort` int NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '任务状态',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统_定时任务' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gf_sys_cron
-- ----------------------------

-- ----------------------------
-- Table structure for gf_sys_cron_group
-- ----------------------------
DROP TABLE IF EXISTS `gf_sys_cron_group`;
CREATE TABLE `gf_sys_cron_group`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务分组ID',
  `pid` bigint NOT NULL COMMENT '父类任务分组ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '分组名称',
  `is_default` tinyint(1) NULL DEFAULT 0 COMMENT '是否默认',
  `sort` int NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '分组状态',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统_定时任务分组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gf_sys_cron_group
-- ----------------------------

-- ----------------------------
-- Table structure for gf_user
-- ----------------------------
DROP TABLE IF EXISTS `gf_user`;
CREATE TABLE `gf_user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `f_id` int NOT NULL,
  `delete_at` int NOT NULL DEFAULT 0,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `infoId` int NULL DEFAULT NULL,
  PRIMARY KEY (`id` DESC) USING BTREE,
  INDEX `f_id`(`f_id` ASC) USING BTREE,
  INDEX `info_id`(`infoId` ASC) USING BTREE,
  CONSTRAINT `f_id` FOREIGN KEY (`f_id`) REFERENCES `gf_famliy` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `info_id` FOREIGN KEY (`infoId`) REFERENCES `gf_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `delete_at` CHECK (`delete_at` in (0,1))
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gf_user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
