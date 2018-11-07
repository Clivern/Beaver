CREATE TABLE IF NOT EXISTS `migration` (
  `id` int(9) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `migration` (`migration`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;