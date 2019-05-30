#!/usr/bin/env bash
#sleep 90s

#run the setup script to create the DB and the schema in the DB
mysql -u docker -pdocker monoDB < "/docker-entrypoint-initdb.d/create-user-table.sql"
mysql -u docker -pdocker monoDB < "/docker-entrypoint-initdb.d/create-tag-table.sql"
mysql -u docker -pdocker monoDB < "/docker-entrypoint-initdb.d/create-data-table.sql"
mysql -u docker -pdocker monoDB < "/docker-entrypoint-initdb.d/create-item-table.sql"
mysql -u docker -pdocker monoDB < "/docker-entrypoint-initdb.d/create-item-data-table.sql"

