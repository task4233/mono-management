-- Token Table
/*
  Depending on User Table
  Necessary to exist User Table
*/

drop table if exists `Token`;

create table if not exists `Token`
(
  `token`          VARCHAR(64) primary key,
  `userId`         INT(8),
  constraint foreignKey_userId_from_Token_to_User
    foreign key (`userId`)
    references `User` (`id`)
    on delete restrict on update restrict
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

