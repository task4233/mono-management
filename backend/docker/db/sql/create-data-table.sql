-- Data Table
drop table if exists `Data`;

create table if not exists `Data`
(
  `id`             INT(8) AUTO_INCREMENT primary key,
  `name`           VARCHAR(64) NOT NULL,
  `type`           VARCHAR(64) NOT NULL,
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

