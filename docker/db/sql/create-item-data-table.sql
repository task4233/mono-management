-- ItemData Table
/*
  Depending on Item Table & Data Table
  Necessary to exist Item Table & Data Table 
*/

drop table if exists `ItemData`;

create table if not exists `ItemData`
(
  `dataId`         INT(8) AUTO_INCREMENT,
  `itemId`         INT(8),
  `num`            double,
  `str`            VARCHAR(64),
  `timestamp`      Date,
  constraint foreignKey_itemId_from_ItemData_to_Item
    foreign key (`itemId`)
    references `Item` (`id`)
    on delete restrict on update restrict,
  constraint foreignKey_dataId_from_ItemData_to_Data
    foreign key (`dataId`)
    references `Data` (`id`)
    on delete restrict on update restrict,
  primary key(`dataId`, `itemId`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

