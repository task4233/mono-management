-- Item Table
drop table if exists `Item`;

create table if not exists `Item`
(
  `id`             INT(8) AUTO_INCREMENT primary key,
  `name`           VARCHAR(64) NOT NULL,
  `userId`         VARCHAR(64) NOT NULL,
  `tagId`          VARCHAR(64) NOT NULL,
  constraint fk_userId
    foreign key (`userId`)
    references User (`id`)
    on delete restrict on update restrict,
  constraint fk_tagId
    foreign key (`tagId`)
    references Tag (`id`)
    on delete restrict on update restrict
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

