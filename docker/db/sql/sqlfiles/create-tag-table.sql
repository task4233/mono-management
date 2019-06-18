-- tags Table
/*
  Depending on users Table & tags Table
  Necessary to exist users Table & tags Table
*/

drop table if exists `tags`;

create table if not exists `tags`
(
  `id`             INT(8) AUTO_INCREMENT primary key,
  `name`           VARCHAR(64) NOT NULL,
  `parentId`       INT(8),
  `userId`         INT(8) NOT NULL,
  constraint foreignKey_parentId_from_tags_to_tags
    foreign key (`parentId`)
    references `tags` (`id`)
    on delete restrict on update restrict,
  constraint foreignKey_userId_from_tags_to_users
    foreign key (`userId`)
    references `users` (`id`)
    on delete restrict on update restrict
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
