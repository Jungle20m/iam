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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


