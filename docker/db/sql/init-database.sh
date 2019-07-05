#!/usr/bin/env bash
#sleep 90s

#run the setup script to create the DB and the schema in the DB
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/sqlfiles/drop-all-tables.sql"
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/sqlfiles/create-user-table.sql"
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/sqlfiles/create-tag-table.sql"
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/sqlfiles/create-data-table.sql"
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/sqlfiles/create-item-table.sql"
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/sqlfiles/create-item-data-table.sql"
mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB < "/docker-entrypoint-initdb.d/sqlfiles/create-token-table.sql"


