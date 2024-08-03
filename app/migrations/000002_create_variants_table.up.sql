CREATE TABLE IF NOT EXISTS variants (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `product_id` int unsigned NOT NULL COMMENT 'プロダクトID',
    `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'バリアント名',
    `price` int NULL COMMENT '価格',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    `deleted_at` datetime DEFAULT null COMMENT '削除日時',
    PRIMARY KEY (`id`)
) COMMENT = 'バリアント';
