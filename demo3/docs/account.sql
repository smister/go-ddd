CREATE TABLE `account` (
    `id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'id主键',
    `amount` decimal(18, 2) NOT NULL DEFAULT '0' COMMENT '价格',
    `province` varchar(50) DEFAULT NULL COMMENT '省份',
    `city` varchar(50) DEFAULT NULL COMMENT '城市',
    `district` varchar(100) DEFAULT NULL COMMENT '行政区',
    `address` varchar(1024) DEFAULT NULL COMMENT '地址',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `integral` (
    `id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'id主键',
    `integral` decimal(18, 2) NOT NULL DEFAULT '0' COMMENT '积分',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `bank_card` (
    `id` varchar(50) DEFAULT NULL COMMENT '银行卡id',
    `account_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '账号ID',
    `bank_name` varchar(50) DEFAULT NULL COMMENT '银行名称',
    `status` tinyint(1) DEFAULT NULL DEFAULT '1' COMMENT '状态',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB;


INSERT INTO `account` (`id`, `amount`) VALUES (1000000000000001, 100),(1000000000000002, 100);
INSERT INTO `integral` (`id`, `integral`) VALUES (1000000000000001, 0),(1000000000000002, 0);