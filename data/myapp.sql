/*
 Navicat Premium Data Transfer

 Source Server         : laragon数据库
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost:3306
 Source Schema         : myapp

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 31/10/2020 17:08:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `me_id` int(11) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_publish` tinyint(4) NULL DEFAULT NULL,
  `created_time` datetime(0) NULL DEFAULT NULL,
  `updated_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, 1, 'Future Communications Officer', 'wyoming_rubber_taiwan.mp4v', 'Doloremque necessitatibus voluptates.\nSequi amet magnam deserunt cupiditate ea sit et aspernatur commodi.\nAliquam vel sit.\nVoluptas eius eveniet ut ea necessitatibus velit qui ducimus.\nNemo recusandae qui esse vitae provident.', 1, '2020-10-31 12:05:55', '2020-10-31 12:05:55');
INSERT INTO `article` VALUES (2, 1, 'Internal Factors Planner', 'gb.shtml', 'sit', 1, '2020-10-31 12:06:53', '2020-10-31 12:06:53');
INSERT INTO `article` VALUES (3, 1, 'Investor Brand Facilitator', 'triple_buffered_avon.jpg', 'reprehenderit et quo', 1, '2020-10-31 12:06:55', '2020-10-31 12:06:55');
INSERT INTO `article` VALUES (4, 1, 'Forward Tactics Liaison', 'frozen_marketing_well_modulated.pdf', 'rerum', 1, '2020-10-31 12:08:28', '2020-10-31 12:08:28');
INSERT INTO `article` VALUES (8, 1, 'Future Marketing Officer', 'jersey.wav', 'odio iusto suscipit', 1, '2020-10-31 13:18:28', '2020-10-31 13:18:28');
INSERT INTO `article` VALUES (9, 1, 'District Implementation Representative', 'assistant.m2a', 'Quos at veniam magni nisi praesentium. Qui nostrum magni aut consequatur necessitatibus dolorum. Iste esse iure optio nulla veritatis quis ipsam qui laborum. Itaque velit repudiandae magnam inventore.', 1, '2020-10-31 13:18:55', '2020-10-31 13:18:55');

-- ----------------------------
-- Table structure for article_tag
-- ----------------------------
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NULL DEFAULT NULL,
  `tag_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article_tag
-- ----------------------------
INSERT INTO `article_tag` VALUES (1, 8, 1);
INSERT INTO `article_tag` VALUES (2, 8, 2);
INSERT INTO `article_tag` VALUES (3, 8, 3);
INSERT INTO `article_tag` VALUES (4, 9, 1);
INSERT INTO `article_tag` VALUES (5, 9, 2);
INSERT INTO `article_tag` VALUES (6, 9, 3);

-- ----------------------------
-- Table structure for me
-- ----------------------------
DROP TABLE IF EXISTS `me`;
CREATE TABLE `me`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `motto` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_time` datetime(0) NULL DEFAULT NULL,
  `updated_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of me
-- ----------------------------
INSERT INTO `me` VALUES (1, 'Jason', '没有文化的人不伤心', '0007', '123456', '2020-10-18 15:22:54', '2020-10-18 15:22:57');

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tag
-- ----------------------------
INSERT INTO `tag` VALUES (1, '小可爱');
INSERT INTO `tag` VALUES (2, '小毛驴');
INSERT INTO `tag` VALUES (3, '小破车');
INSERT INTO `tag` VALUES (4, '小红毛');

SET FOREIGN_KEY_CHECKS = 1;
