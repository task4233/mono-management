# Overview
 - `frontend`にはフロントのものを入れてください
 - `backend`にはバックのものを入れてください
 
 # In Charge
 ## Webデザイナー
  - mocchey (:vue:)
 ## フロントエンジニア(:vue:)
  - MissionaryPS
  - shirakawa36
 ## バックエンジニア(:gopher_dance:)
  - cordx56
  - task4233
  - ka-tsu-mo

# Run
## Mysql

```bash
$ docker-compose up -d

# 外部からMySQLにつなぐ場合は以下の通り
$  mysql --defaults-extra-file=./docker/db/my.cnf -h 0.0.0.0 -P 3306 --protocol=tcp monoDB

# 直接MySQLをいじりたい場合は以下の通り
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                               NAMES
da1c2ef99131        mysql:8.0           "docker-entrypoint.s…"   3 seconds ago       Up 2 seconds        0.0.0.0:3306->3306/tcp, 33060/tcp   mysql_host
$ docker exec -it da1 /bin/bash
$ mysql --defaults-extra-file=/etc/mysql/my.cnf monoDB

# 落とすとき
$ docker-compose down
```
