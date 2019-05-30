-- ItemData Table
drop table if exists `ItemData`;

create table if not exists `ItemData`
(
  `dataId`         INT(8) AUTO_INCREMENT primary key,
  `itemId`         INT(8) primary key,
  `num`            double,
  `str`            VARCHAR(64),
  `timestamp`      Date,
  constraint fk_itemId
    foreign key (`itemId`)
    references `Item` (`id`)
    on delete restrict on update restrict,
  constraint fk_dataId
    foreign key (`dataId`)
    references (`Data`)
    on delete restrict on update restrict
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

