postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres dropdb --username=postgres simple_bank

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up

.PHONY: postgres createdb dropdb migratedown migrateup