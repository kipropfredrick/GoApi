postgresinit:
	docker run --name postgres15 -p 5433:5433 -e POSTGRES_USER=kiprop -e POSTGRES_PASSWORD=31877101 -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=kiprop --owner=kiprop go-chat

dropdb:
	docker exec -it postgres15 dropdb go-chat

migrateup:
	migrate -path db/migrations -database "postgresql://kiprop:31877101@kiprop.cs7pago4pgpd.us-west-2.rds.amazonaws.com:5432/gochat?sslmode=disable"  up

migratedown:
	migrate -path db/migrations -database "postgresql://kiprop:31877101@kiprop.cs7pago4pgpd.us-west-2.rds.amazonaws.com:5432/gochat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown

# migrate create -ext sql -dir db/migrations/ -seq account_validation