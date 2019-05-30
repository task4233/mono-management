-- User Table
drop table if exists `User`;

create table if not exists `User`
(
  `id`             INT(8) AUTO_INCREMENT primary key,
  `name`           VARCHAR(64) NOT NULL,
  `hashed_pass`    VARCHAR(64) NOT NULL,
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

