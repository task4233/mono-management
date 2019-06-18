-- tokens Table
/*
  Depending on users Table
  Necessary to exist users Table
*/

drop table if exists `tokens`;

create table if not exists `tokens`
(
  `token`          VARCHAR(64) primary key,
  `userId`         INT(8) NOT NULL,
  constraint foreignKey_userId_from_tokens_to_users
    foreign key (`userId`)
    references `users` (`id`)
    on delete restrict on update restrict
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
