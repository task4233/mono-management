-- users Table
drop table if exists `users`;

create table if not exists `users`
(
  `id`             INT(8) AUTO_INCREMENT,
  `name`           VARCHAR(64) NOT NULL,
  `hashed_pass`    VARCHAR(64) NOT NULL,
  primary key(`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

