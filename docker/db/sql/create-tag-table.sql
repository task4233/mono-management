-- Tag Table
drop table if exists `Tag`;

create table if not exists `Tag`
(
  `id`             INT(8) AUTO_INCREMENT primary key,
  `name`           VARCHAR(64) NOT NULL,
  `parentId`       INT(8),
  `userId`         INT(8),
  constraint fk_parentId
    foreign key (`parentId`)
    references `Tag` (`id`)
    on delete restrict on update restrict,
  constraint fk_userId
    foreign key (`userId`)
    references `User` (`id`)
    on delete restrict on update restrict
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

