
GoSimpleBank>docker pull postgres:14.5-alpine

GoSimpleBank>docker images

GoSimpleBank>docker run --name postgres14.5 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretkey -d postgres:14.5-alpine

GoSimpleBank>docker ps

GoSimpleBank>docker exec -it postgres14.5 psql -U root


GoSimpleBank>docker logs postgres14.5

//spin the docker
Gosimplebank>docker compose up -d

Gosimplebank>docker exec -it postgres14.5 /bin/sh

># createdb --username=root --owner=root simple_bank
># psql simple_bank
># dropdb simple_bank
># exit

Gosimplebank>docker exec -it postgres14.5 createdb --username=root --owner=root simple_bank

Gosimplebank>docker exec -it postgres14.5 psql -U root simple_bank

GoSimpleBank>migrate create -ext sql -dir db/migration -seq init_schema
GoSimpleBank>migrate create -ext sql -dir db/migration -seq add_users
GoSimpleBank>migrate create -ext sql -dir db/migration -seq add_sessions
GoSimpleBank>migrate create -ext sql -dir db/migration -seq add_verify_emails
GoSimpleBank>migrate create -ext sql -dir db/migration -seq add_role_to_users

GoSimpleBank>make postgres

GoSimpleBank>make createdb

GoSimpleBank>migrate -path db/migration -database "postgresql://root:secretkey@localhost:5432/simple_bank?sslmode=disable" -verbose up

Gosimplebank>go mod tidy

Gosimplebank>


https://github.com/kyleconroy/sqlc


---------------------------------------
//create project
CSharp>dotnet new webapi -n RedisAPI

//spin the docker
RedisAPI>docker compose up -d

//list of docker
RedisAPI>docker ps

-------------------------------------

RedisAPI>docker exec -it <container_id> /bin/bash
example: docker exec -it 3f9bdbdf1c1f /bin/bash

//learn about key value from redis
//https://redis.io/commands
root@3f9bdbdf1c1f:/data# redis-cli
127.0.0.1:6379> ping 
PONG
127.0.0.1:6379>set platform:10010 Docker
OK
127.0.0.1:6379>get platform:10010
"Docker"
127.0.0.1:6379>del platform:10010
(integer) 1
127.0.0.1:6379>scan 0
give result
127.0.0.1:6379>exit
root@3f9bdbdf1c1f:/data#exit
RedisAPI>
-------------------------------------
Which Package?
1. Microsoft.Extensions.Caching.Redis(deprecated)
2. Microsoft.Extensions.Caching.StackExchangeRedis (IDistributedCache & IConnectionMultiplexer)
3. StackExchange.Redis (IConnectionMultiplexer)

Boils Down to whether you want to use:
- IDistributedCache (restricted datatypes)
- IConnectionMultiplexer (All datatypes available)

RedisAPI>dotnet add package Microsoft.Extensions.Caching.StackExchangeRedis --version 6.0.2
-------------------------------------

End Points
Description                 Verb    Route
Retrive a Single resource   GET     api/platforms/{id}
Create a resource           POST    api/platforms
Retrive All resource        GET     api/platforms
-------------------------------------

POST Request:
https://localhost:7002/api/platforms
GET Request:
https://localhost:7002/api/platforms?platform=25be7e3a-2962-4578-af1c-6c6110d06f30

-------------------------------

Sets
-Unordered collections of strings
-1 to many mapping between key & value
-set using SADD
    - sadd myset 1 2 3
- Returned using SMEMBERS
    -smembers myset
-Check if value present in set SISMEMBER


------------------
Hashes
-stores field/value pairs
- suitable for storing structure objects
- set using HMSET
    -hmset <id> <field1> <value1>
- get all items using HGETALL
    - hgetall <id>
- get individual items using HGET
    - hget <id> <field1>

