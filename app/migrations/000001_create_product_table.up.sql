CREATE TABLE IF NOT EXISTS products (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品名',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    `deleted_at` datetime DEFAULT null COMMENT '削除日時',
    PRIMARY KEY (`id`)
);
