-- ----------------------------
--  Table structure for `users`
-- ----------------------------
CREATE TABLE `users`
(
    `uid`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `email`         varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱账户',
    `password`      char(60)     NOT NULL DEFAULT '' COMMENT '密码',
    `nickname`      varchar(50)  NOT NULL COMMENT '用户昵称',
    `real_name`     varchar(30)  NOT NULL DEFAULT '' COMMENT '姓名',
    `is_enable`     tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '用户启用状态',
    `register_ip`   varchar(15)  NOT NULL DEFAULT '' COMMENT '注册或添加IP',
    `register_at`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册或添加时间',
    `login_times`   int(10) unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
    `last_login_ip` varchar(15)  NOT NULL DEFAULT '' COMMENT '最后登录IP',
    `last_login_at` datetime     NOT NULL DEFAULT '1000-01-01 01:01:01' COMMENT '最后登录时间',
    `updated_at`    datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后数据更新时间',
    PRIMARY KEY (`uid`),
    UNIQUE KEY `uk_nickname` (`nickname`),
    UNIQUE KEY `uk_email` (`email`),
    KEY             `idx_realName` (`real_name`)
) ENGINE=InnoDB AUTO_INCREMENT=100000001 DEFAULT CHARSET=utf8 COMMENT='用户';

