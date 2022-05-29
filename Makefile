postgres:
	docker run --name postgres13 -p 5342:5432 -e POSTGRES_USER=postgres12 -e POSTGRES_PASSWORD=secret -d postgres:14.3-alpine

createdb:
	docker exec -it postgres13 createdb --username=postgres12 --owner=postgres12 simple_bank
	
dropdb:
	docker exec -it postgres13 dropdb --username=postgres12 simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres12:secret@localhost:5342/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres12:secret@localhost:5342/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test