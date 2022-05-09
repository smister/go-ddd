CREATE TABLE `account` (
    `id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'id主键',
    `amount` decimal(18, 2) NOT NULL DEFAULT '0' COMMENT '价格',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB;

INSERT INTO `account` (`id`, `amount`) VALUES (1000000000000001, 100),(1000000000000002, 100);