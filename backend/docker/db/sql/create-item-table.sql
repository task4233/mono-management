-- Item Table
drop table if exists `Item`;

create table if not exists `Item`
(
  `id`             INT(8) AUTO_INCREMENT,
  `name`           VARCHAR(64) NOT NULL,
  `userId`         INT(8) NOT NULL,
  `tagId`          INT(8) NOT NULL,
  constraint fk_userId
    foreign key (`userId`)
    references User (`id`)
    on delete restrict on update restrict,
  constraint fk_tagId
    foreign key (`tagId`)
    references Tag (`id`)
    on delete restrict on update restrict,
    primary key(`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

