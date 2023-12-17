
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


