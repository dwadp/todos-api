CREATE TABLE IF NOT EXISTS `todos` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` longtext,
  `description` longtext,
  `is_done` tinyint(1) DEFAULT NULL,
  `due_date` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;