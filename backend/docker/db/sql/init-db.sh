#!/usr/bin/env bash
#sleep 90s

mysql -u docker -pdocker monoDB < "/docker-entrypoint-initdb.d/001-create-tables.sql"
