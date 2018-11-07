CREATE TABLE IF NOT EXISTS `config` (
  `id` int(9) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(60) NOT NULL,
  `value` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;