
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

GoSimpleBank>make migratedown
GoSimpleBank>make migrateup
GoSimpleBank>make postgres

GoSimpleBank>make createdb

GoSimpleBank>migrate -path db/migration -database "postgresql://root:secretkey@localhost:5432/simple_bank?sslmode=disable" -verbose up

Gosimplebank>go mod tidy

Gosimplebank>


https://github.com/kyleconroy/sqlc


Install gin framework
https://github.com/gin-gonic/gin

Gosimplebank>go get -u github.com/gin-gonic/gin

----

Viper: https://github.com/spf13/viper

GoSimpleBank>go get github.com/spf13/viper

----
gomock

https://github.com/golang/mock

GoSimpleBank>go install github.com/golang/mock/mockgen@v1.6.0
GoSimpleBank>which mockgen
GoSimpleBank>vi ~/.bash_profile //add export PATH=$PATH:~/go/bin
GoSimpleBank>source ~/.bash_profile  //reload bash_profile
GoSimpleBank>which mockgen

GoSimpleBank>mockgen -destination destination_path Import_path_store_interface InterfaceName
GoSimpleBank>mockgen -destination db/mock/store.go gosimplebank/db/sqlc Store
//change package name
GoSimpleBank>mockgen -package mockdb -destination db/mock/store.go gosimplebank/db/sqlc Store
// add mock to makefile

GoSimpleBank>git add .
GoSimpleBank>git commit -m "Corrected test run"
GoSimpleBank>git push -f origin main 

----------------------------------------------------------------
GoSimpleBank>go get github.com/google/uuid
GoSimpleBank>go get github.com/golang-jwt/jwt/v4
GoSimpleBank>go get github.com/o1egl/paseto
GoSimpleBank>go get github.com/aead/chacha20poly1305
----------------------------------------------------------------
//Implement authentication middleware and authorization rules

----------------------------------------------------------------
#23 Build a minimal Golang Docker image with a multistage Dockerfile
// move from main git branch to another working branch
GoSimpleBank>git checkout -b ft/docker
//check current status
GoSimpleBank>git status

//build docker image
GoSimpleBank>docker build -t gosimplebank:latest .
//list docker images
GoSimpleBank>docker images

----------------------------------------------------------------
#24 How to use docker network to connect 2 stand-alone containers
GoSimpleBank>docker rm gosimplebanks
GoSimpleBank>docker ps -a
GoSimpleBank>docker rm 58fsdf332(imageid)

GoSimpleBank>docker run --name gosimplebank -p 8080:8080 gosimplebanks:latest
//Run GIN in release mode 
GoSimpleBank>docker run --name gosimplebank -p 8080:8080 -e GIN_MODE=release gosimplebanks:latest
// to see postgres container network settings
GoSimpleBank>docker container inspect postgres14.5
GoSimpleBank>docker container inspect gosimplebank

//easiest way to fix postgres network to connect update app.env, but not good every time need to build image, use viper
GoSimpleBank>docker run --name gosimplebank -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secretkey@172.17.0.3:5432/simple_bank?sslmode=disable" gosimplebanks:latest

GoSimpleBank>docker network ls

GoSimpleBank>docker network inspect bridge(i.e. network name)

GoSimpleBank>docker network create bank-network
GoSimpleBank>docker network connect bank-network postgres14.5
GoSimpleBank>docker network inspect bank-network
GoSimpleBank>docker network inspect postgres14.5

// to put app & postgres in same network
GoSimpleBank>docker run --name gosimplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secretkey@postgres14.5:5432/simple_bank?sslmode=disable" gosimplebanks:latest

GoSimpleBank>git add .
GoSimpleBank>git commit -m "add docker file"
GoSimpleBank>git push origin ft/docker

----------------------------------------------------------------
#25 How to write docker-compose file and control service start-up orders with wait-for.sh
https://docs.docker.com/compose/compose-file/compose-file-v3/
GoSimpleBank>git checkout main
GoSimpleBank>git pull
GoSimpleBank>docker compose up
GoSimpleBank>chmod +x start.sh

GoSimpleBank>docker compose down
GoSimpleBank>docker rmi gosimplebanks_api
GoSimpleBank>docker compose up

//add wait-for.sh
//https://github.com/eficode/wait-for
GoSimpleBank>chmod +x wait-for.sh
GoSimpleBank>docker compose down
GoSimpleBank>docker rmi gosimplebanks_api
GoSimpleBank>docker compose up

