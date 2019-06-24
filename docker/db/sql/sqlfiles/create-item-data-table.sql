-- itemdatas Table
/*
  Depending on items Table & datas Table
  Necessary to exist items Table & datas Table 
*/

drop table if exists `itemdatas`;

create table if not exists `itemdatas`
(
  `dataId`         INT(8) AUTO_INCREMENT,
  `itemId`         INT(8),
  `num`            double,
  `str`            VARCHAR(64),
  `timestamp`      Date,
  constraint foreignKey_itemId_from_itemdatas_to_items
    foreign key (`itemId`)
    references `items` (`id`)
    on delete cascade on update restrict,
  constraint foreignKey_dataId_from_itemdatas_to_datas
    foreign key (`dataId`)
    references `datas` (`id`)
    on delete cascade on update restrict,
  primary key(`dataId`, `itemId`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
