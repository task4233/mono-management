-- Item Table
/*
  Depending on User Table & Tag Table
  Necessary to exist User Table & Tag Table
*/

drop table if exists `Item`;

create table if not exists `Item`
(
  `id`             INT(8) AUTO_INCREMENT,
  `name`           VARCHAR(64) NOT NULL,
  `userId`         INT(8) NOT NULL,
  `tagId`          INT(8) NOT NULL,
  constraint foreignKey_userId_from_Item_to_User
    foreign key (`userId`)
    references User (`id`)
    on delete restrict on update restrict,
  constraint foreignKey_tagId_from_Item_to_Tag
    foreign key (`tagId`)
    references Tag (`id`)
    on delete restrict on update restrict,
    primary key(`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

