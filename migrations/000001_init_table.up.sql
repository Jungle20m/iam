CREATE TABLE `healthnet`.`one_time_password` (
     `id` int NOT NULL AUTO_INCREMENT,
     `user_id` INT NOT NULL,
     `client_id` VARCHAR(255) NULL,
     `phone_number` VARCHAR(15) NULL,
     `message_body` TEXT NULL,
     `otp` varchar(45) DEFAULT NULL,
     `expired` int DEFAULT NULL,
     `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
     `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `healthnet`.`user_account` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_name` varchar(45) DEFAULT NULL,
    `phone_number` varchar(15) DEFAULT NULL,
    `email` varchar(45) DEFAULT NULL,
    `password` varchar(255) DEFAULT NULL,
    `password_salt` varchar(255) DEFAULT NULL,
    `password_hash_algorithms` varchar(45) DEFAULT NULL,
    `user_status` varchar(45) DEFAULT NULL,
    `registration_time` datetime DEFAULT NULL,
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `healthnet`.`user_token` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `client_id` varchar(255) NULL,
    `id_token` TEXT NOT NULL,
    `access_token` TEXT NOT NULL,
    `refresh_token` TEXT NOT NULL,
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



