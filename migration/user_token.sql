CREATE TABLE `healthnet`.`user_token` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `client_id` varchar(255) NULL,
    `id_token` varchar(255) NOT NULL,
    `access_token` varchar(255) NOT NULL,
    `refresh_token` varchar(255) NOT NULL,
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


