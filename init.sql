CREATE DATABASE IF NOT EXISTS `onycom` /*!40100 COLLATE 'utf8mb4_general_ci' */;

USE `onycom`;

CREATE TABLE `user` (
	`id` VARCHAR(11) NOT NULL COLLATE 'utf8mb4_general_ci',
	`created_at` TIMESTAMP NOT NULL DEFAULT current_timestamp(),
	`email` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
	`password` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
	`name` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `email` (`email`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;

CREATE TABLE `post` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NOT NULL DEFAULT current_timestamp(),
	`updated_at` TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	`user_id` VARCHAR(11) NOT NULL COLLATE 'utf8mb4_general_ci',
	`title` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
	`content` TEXT NOT NULL COLLATE 'utf8mb4_general_ci',
	`view_count` INT(11) NOT NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `fk_user` (`user_id`) USING BTREE,
	CONSTRAINT `fk_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=18
;
