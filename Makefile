postgres:
	docker-compose up -d

createdb:
	docker exec -it golang_postgres createdb --username=postgres --owner=postgres simple_stores

dropdb:
	docker exec -it golang_postgres dropdb --username=postgres simple_stores

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/simple_stores?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/simple_stores?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

install:
	go mod tidy

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/thuhangpham/simplestores/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock