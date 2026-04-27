CREATE TABLE IF NOT EXISTS `user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    `passport` varchar(45) NOT NULL COMMENT 'User Passport',
    `nickname` varchar(45) NOT NULL COMMENT 'User Nickname',
    `password` varchar(45) NOT NULL COMMENT 'User Password',
    `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
    `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
    `delete_at` datetime DEFAULT NULL COMMENT 'Deleted Time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_passport` (`passport`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
