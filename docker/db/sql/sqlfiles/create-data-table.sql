-- datas Table

drop table if exists `datas`;

create table if not exists `datas`
(
  `id`             INT(8) AUTO_INCREMENT primary key,
  `name`           VARCHAR(64) NOT NULL,
  `type`           VARCHAR(64) NOT NULL
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

