CREATE TABLE `users` (
     `id` int NOT NULL AUTO_INCREMENT,
     `name` varchar(45) DEFAULT NULL,
     `phone` varchar(45) DEFAULT NULL,
     `email` varchar(45) DEFAULT NULL,
     `password` varchar(255) NOT NULL,
     `is_active` tinyint DEFAULT NULL,
     `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
     `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci