create database if not exists `bilibili_video` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';

USE bilibili_video;


DROP TABLE IF EXISTS `t_video`;
CREATE TABLE `t_video` (
                           `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
                           `user_id` bigint NOT NULL COMMENT '用户id',
                           `url` varchar(500) NOT NULL COMMENT '视频链接',
                           `thumbnail` varchar(500) NOT NULL COMMENT '封面链接',
                           `title` varchar(255) NOT NULL COMMENT '视频标题',
                           `types` int NOT NULL COMMENT '视频类型：0原创，1转载',
                           `duration` varchar(255) NOT NULL COMMENT '视频时长',
                           `area` varchar(255) NOT NULL COMMENT '所在分区:0鬼畜，1音乐，2电影',
                           `description` text COMMENT '视频简介',
                           `created_at` datetime(3) NULL COMMENT '创建时间',
                           `updated_at` datetime(3) NULL COMMENT '更新时间',
                           `deleted_at` datetime(3) NULL COMMENT '删除时间',
                           PRIMARY KEY (`id`),INDEX `idx_t_video_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频投稿记录表';

DROP TABLE IF EXISTS `t_video_coin`;
CREATE TABLE `t_video_coin` (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
                                `user_id` bigint DEFAULT NULL COMMENT '用户id',
                                `video_id` bigint DEFAULT NULL COMMENT '视频投稿id',
                                `amount` int DEFAULT NULL COMMENT '投币数',
                                `created_at` datetime(3) NULL COMMENT '创建时间',
                                `updated_at` datetime(3) NULL COMMENT '更新时间',
                                `deleted_at` datetime(3) NULL COMMENT '删除时间',
                                PRIMARY KEY (`id`),INDEX `idx_t_video_coin_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频投币记录表';

DROP TABLE IF EXISTS `t_video_collection`;
CREATE TABLE `t_video_collection` (
                                      `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
                                      `video_id` bigint DEFAULT NULL COMMENT '视频投稿id',
                                      `user_id` bigint DEFAULT NULL COMMENT '用户id',
                                      `group_id` bigint DEFAULT NULL COMMENT '收藏分组',
                                      `created_at` datetime(3) NULL COMMENT '创建时间',
                                      `updated_at` datetime(3) NULL COMMENT '更新时间',
                                      `deleted_at` datetime(3) NULL COMMENT '删除时间',
                                      PRIMARY KEY (`id`),INDEX `idx_t_video_collection_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频收藏表';

DROP TABLE IF EXISTS `t_video_comment`;
CREATE TABLE `t_video_comment` (
                                   `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
                                   `video_id` bigint NOT NULL COMMENT '视频id',
                                   `user_id` bigint NOT NULL COMMENT '用户id',
                                   `comment` text NOT NULL COMMENT '评论',
                                   `reply_user_id` bigint DEFAULT NULL COMMENT '回复用户id',
                                   `root_id` bigint DEFAULT NULL COMMENT '根结点评论id',
                                   `created_at` datetime(3) NULL COMMENT '创建时间',
                                   `updated_at` datetime(3) NULL COMMENT '更新时间',
                                   `deleted_at` datetime(3) NULL COMMENT '删除时间',
                                   PRIMARY KEY (`id`),INDEX `idx_t_video_comment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频评论表';

DROP TABLE IF EXISTS `t_video_like`;
CREATE TABLE `t_video_like` (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
                                `user_id` bigint NOT NULL COMMENT '用户id',
                                `video_id` bigint NOT NULL COMMENT '视频投稿id',
                                `created_at` datetime(3) NULL COMMENT '创建时间',
                                `updated_at` datetime(3) NULL COMMENT '更新时间',
                                `deleted_at` datetime(3) NULL COMMENT '删除时间',
                                PRIMARY KEY (`id`),INDEX `idx_t_video_like_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频点赞表';

DROP TABLE IF EXISTS `t_barrage`;
CREATE TABLE `t_barrage` (
                           `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
                           `user_id` bigint DEFAULT NULL COMMENT '用户id',
                           `video_id` bigint DEFAULT NULL COMMENT '视频id',
                           `content` text COMMENT '弹幕内容',
                           `barrage_time` int DEFAULT NULL COMMENT '弹幕出现时间',
                           `created_at` datetime(3) NULL COMMENT '创建时间',
                           `updated_at` datetime(3) NULL COMMENT '更新时间',
                           `deleted_at` datetime(3) NULL COMMENT '删除时间',
                           PRIMARY KEY (`id`),INDEX `idx_t_video_like_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='弹幕表';

