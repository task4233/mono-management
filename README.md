# Overview
 - `frontend`にはフロントのものを入れてください
 - `backend`にはバックのものを入れてください
 
 # In Charge
 ## Webデザイナー
  - @持田達也 (:vue:)
 ## フロントエンジニア(:vue:)
  - @al17001_赤松 寛基
  - @唐沢将太_AL17037
 ## バックエンジニア(:gopher_dance:)
  - @al17029_岡本祐希
  - @al17111_美馬隆志
  - @al17036_加茂司

# Run
## Mysql

```bash
$ docker-compose up -d

# 直接MySQLをいじりたい場合は以下の通り
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                               NAMES
da1c2ef99131        mysql:8.0           "docker-entrypoint.s…"   3 seconds ago       Up 2 seconds        0.0.0.0:3306->3306/tcp, 33060/tcp   mysql_host
$ docker exec -it da1 /bin/bash

# 落とすとき
$ docker-compose down
```
