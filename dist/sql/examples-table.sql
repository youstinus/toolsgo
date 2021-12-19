CREATE TABLE `examples` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `example` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `example` (`example`),
) ENGINE=InnoDB;
