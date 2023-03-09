.PHONY: createdb dropdb migrateup migratedown sqlc

createdb:
	docker exec -it postgres15 createdb --username=Reoptima --owner=Reoptima simple_app

dropdb:
	docker exec -it postgres15 dropdb simple_app

migrateup:
	migrate -path db/migration -database "postgresql://Reoptima:passwd@localhost:5432/simple_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://Reoptima:passwd@localhost:5432/simple_app?sslmode=disable" -verbose down
sqlc:
	docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate
test:
	go test -v -cover ./...