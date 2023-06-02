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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
