-- Token Table
drop table if exists `Token`;

create table if not exists `Token`
(
  `token`          VARCHAR(64) primary key,
  `userId`         INT(8),
  constraint fk_userId
    foreign key (`userId`)
    references `User` (`id`)
    on delete restrict on update restrict
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

