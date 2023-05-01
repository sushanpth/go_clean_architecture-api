
-- +migrate Up
CREATE TABLE IF NOT EXISTS `businesses` (
  `id` BINARY(16) NOT NULL,
  `user_id` BINARY(16) NOT NULL,
  `name` VARCHAR(250) NOT NULL,
  `location` VARCHAR(250) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT FK_User_Business FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;


-- +migrate Down
DROP TABLE IF EXISTS `businesses`;
