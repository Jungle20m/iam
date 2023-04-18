CREATE TABLE `user_verification` (
     `id` int NOT NULL AUTO_INCREMENT,
     `token` varchar(45) DEFAULT NULL,
     `expire_time` datetime DEFAULT NULL,
     `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
     `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci