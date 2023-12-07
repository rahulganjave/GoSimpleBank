
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


Install gin framework
https://github.com/gin-gonic/gin

Gosimplebank>go get -u github.com/gin-gonic/gin