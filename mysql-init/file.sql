create database if not exists `bilibili_file` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';

USE bilibili_file

DROP TABLE IF EXISTS `t_file`;
CREATE TABLE `t_file` (
                          `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
                          `url` varchar(500) DEFAULT NULL COMMENT '文件存储路径',
                          `type` varchar(50) DEFAULT NULL COMMENT '文件类型',
                          `md5` varchar(500) DEFAULT NULL COMMENT '文件MD5唯一标识',
                          `uuid` varchar(500) DEFAULT NULL COMMENT '文件唯一标识',
                          `upload_id` varchar(500) DEFAULT NULL COMMENT 'minio上传id',
                          `created_at` datetime(3) NULL COMMENT '创建时间',
                          `updated_at` datetime(3) NULL COMMENT '更新时间',
                          `deleted_at` datetime(3) NULL COMMENT '删除时间',
                          PRIMARY KEY (`id`), INDEX `idx_t_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='上传文件信息表';