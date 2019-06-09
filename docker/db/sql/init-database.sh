#!/usr/bin/env bash
#sleep 90s

#run the setup script to create the DB and the schema in the DB
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/drop-all-tables.sql"
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/create-user-table.sql"
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/create-tag-table.sql"
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/create-data-table.sql"
# mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/create-item-table.sql"
# mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/create-item-data-table.sql"
# mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/create-token-table.sql"

