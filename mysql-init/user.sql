create database if not exists `bilibili_user` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';

USE bilibili_user

DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
                          `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
                          `username` varchar(100) NOT NULL COMMENT '用户名',
                          `phone` varchar(100) DEFAULT NULL COMMENT '手机号',
                          `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
                          `password` varchar(255) DEFAULT NULL COMMENT '密码',
                          `created_at` datetime(3) NULL COMMENT '创建时间',
                          `updated_at` datetime(3) NULL COMMENT '更新时间',
                          `deleted_at` datetime(3) NULL COMMENT '删除时间',
                          PRIMARY KEY (`id`), INDEX `idx_t_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

DROP TABLE IF EXISTS `t_user_info`;
CREATE TABLE `t_user_info` (
                               `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
                               `user_id` bigint DEFAULT NULL COMMENT '用户id（关联）',
                               `nick` varchar(100) DEFAULT NULL COMMENT '昵称',
                               `avatar` varchar(1024) DEFAULT NULL COMMENT '头像',
                               `sign` text COMMENT '签名',
                               `gender` integer DEFAULT NULL COMMENT '性别：0男，1女，2未知',
                               `birth` varchar(20) DEFAULT NULL COMMENT '生日',
                               `created_at` datetime(3) NULL COMMENT '创建时间',
                               `updated_at` datetime(3) NULL COMMENT '更新时间',
                               `deleted_at` datetime(3) NULL COMMENT '删除时间',
                               PRIMARY KEY (`id`), INDEX `idx_t_user_info_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户基本信息表';

