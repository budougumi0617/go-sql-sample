-- Sample migration

-- +migrate Up
CREATE TABLE `user` (
`id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
`name` varchar(255) NOT NULL COMMENT 'user name',
`email` varchar(255) NOT NULL COMMENT 'e-mail address',
`created_at` datetime,
`updated_at` datetime)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC
COMMENT='user table is the sample table'
;

-- +migrate Down
DROP TABLE IF EXISTS `user`;
