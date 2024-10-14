-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_crm_user_c` (
    `usr_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `usr_email` varchar(30) NOT NULL DEFAULT '',
    `usr_phone` varchar(15) NOT NULL DEFAULT '',
    `usr_username` varchar(30) NOT NULL DEFAULT '',
    `usr_password` varchar(32) NOT NULL DEFAULT '',
    `usr_created_at` int(11) NOT NULL DEFAULT '0',
    `usr_updated_at` int(11) NOT NULL DEFAULT '0',
    `usr_create_ip_at` varchar(12) NOT NULL DEFAULT '',
    `usr_last_login_at` int(11) NOT NULL DEFAULT '0',
    `usr_last_login_ip_at` varchar(12) NOT NULL DEFAULT '',
    `usr_last_login_times` int(11) NOT NULL DEFAULT '0',
    `usr_status` tinyint(11) NOT NULL DEFAULT '0' COMMENT 'Status 1:enable, 0:disable, -1:deleted',
    PRIMARY KEY (`usr_id`),
    KEY `idx_email` (`usr_email`),
    KEY `idx_phone` (`usr_phone`),
    KEY `idx_username` (`usr_username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Account';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS 'pre_go_crm_user_c';
-- +goose StatementEnd
