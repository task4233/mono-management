#!/bin/sh
docker-compose exec db bash -c "chmod 0775 docker-entrypoint-initdb.d/init-db.sh"
