-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_go_acc_user_info_9999` (
    `user_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_account` varchar(255) NOT NULL,
    `user_nickname` varchar(255) DEFAULT NULL,
    `user_avatar` varchar(255) DEFAULT NULL,
    `user_state` tinyint unsigned NOT NULL COMMENT '0-Locked, 1-Activated, 2-Not Activated',
    `user_mobile` varchar(20) DEFAULT NULL,
    `user_gender` tinyint unsigned DEFAULT NULL COMMENT '0-Secret, 1-Male, 2-Female',
    `user_birthday` date DEFAULT NULL,
    `user_email` varchar(255) DEFAULT NULL,
    `user_is_authentication` tinyint unsigned NOT NULL COMMENT '0-Not Authenticated, 1-Authenticated',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `unique_user_account` (`user_account`),
    KEY `idx_user_mobile` (`user_mobile`),
    KEY `idx_user_email` (`user_email`),
    KEY `idx_user_state` (`user_state`),
    KEY `idx_user_is_authentication` (`user_is_authentication`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_info_9999`;
-- +goose StatementEnd
