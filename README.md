## Vision
欲しいなら作ればいいじゃん
## Overview
### Purpose
犬が同伴できるお店を手軽に探す
### Current Problem
- 犬が同伴することを前提とした飲食店検索サービスが少ない
- 犬に寄り添った店舗情報を探せない
- 犬がいると、例えば公園でテイクアウトして食べることしかできない
### Approach
犬同伴時の飲食店の条件をユーザーに情報提供する
### List
- 現在地から(条件)検索
- 地図表示・リスト表示
- 店舗概要
- クチコミ・写真表示
- 後で見る登録
- サインアップ・サインイン
### Image
(暫定案)

![image](https://user-images.githubusercontent.com/64164948/122626275-93756f80-d0e4-11eb-894a-2ecb50c53a49.png)
![image](https://user-images.githubusercontent.com/64164948/122626281-9e300480-d0e4-11eb-8856-53e3803a1a11.png)
![image](https://user-images.githubusercontent.com/64164948/122626292-ac7e2080-d0e4-11eb-8cb2-7dadd80ec4b3.png)
![image](https://user-images.githubusercontent.com/64164948/122626301-b56ef200-d0e4-11eb-82d7-5eb00575e4b5.png)
![image](https://user-images.githubusercontent.com/64164948/122626310-bef85a00-d0e4-11eb-962b-539f4a2967a2.png)
![image](https://user-images.githubusercontent.com/64164948/122626315-c61f6800-d0e4-11eb-9318-ad5f17b4a901.png)

### Architecture
- nginx
- Go
- MySQL replication
- Redis replication
- DynamoDB
- grafana
- kibana
- elasticsearch
- Swagger
### Install Guide
```
$ git clone https://github.com/satorunooshie/eat-with-dog-API.git /go/src/

$ cd eat-with-dog-API && cp .env.example .env

$ docker-compose up -d
```

#### Replication
TODO: Shellに変更
```
$ docker network ls

# copy mysql_slave NETWORK ID

$ docker inspect {NETWORK ID}

# copy mysql_master.IPv4Address
$docker-compose exec mysql-slave bash

$ mysql -u root -p

mysql> SET GLOBAL SQL_SLAVE_SKIP_COUNTER=1;

mysql> CHANGE MASTER TO MASTER_HOST='{mysql_master.IPv4Address}', MASTER_USER='{MYSQL_MASTER_USER}', MASTER_PASSWORD='{MYSQL_MASTER_PASSWORD}', MASTER_LOG_FILE='mysql-bin.000001', MASTER_LOG_POS=0;

mysql> START SLAVE;
```

##### Replication Check
```
mysql> SHOW SLAVE STATUS;
Slave_IO_Running: Yes
Slave_SQL_Runnning: Yes

$ docker-compose exec mysql-master bash

$ mysql -u root -p

mysql> SHOW MASTER STATUS;
```

### Add Endpoint
MUST UPDATE Swagger
### Test For API
```$ docker-compose up -d```
### See Also
