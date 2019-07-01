-- items Table
/*
  Depending on users Table & tags Table
  Necessary to exist users Table & tags Table
*/

drop table if exists `items`;

create table if not exists `items`
(
  `id`             INT(8) AUTO_INCREMENT,
  `name`           VARCHAR(64) NOT NULL,
  `userId`         INT(8) NOT NULL,
  `tagId`          INT(8) NOT NULL,
  constraint foreignKey_userId_from_items_to_users
    foreign key (`userId`)
    references `users` (`id`)
    on delete cascade on update restrict,
  constraint foreignKey_tagId_from_items_to_tags
    foreign key (`tagId`)
    references `tags` (`id`)
    on delete cascade on update restrict,
    primary key(`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

