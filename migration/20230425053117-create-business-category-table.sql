
-- +migrate Up
CREATE TABLE IF NOT EXISTS `business_categories` (
  `id` BINARY(16) NOT NULL,
  `business_id` BINARY(16) NOT NULL,
  `category_id` BINARY(16) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT FK_Business_Category_Business FOREIGN KEY (`business_id`) REFERENCES `businesses`(`id`),
  CONSTRAINT FK_Business_Category_Category FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS `business_categories`;
