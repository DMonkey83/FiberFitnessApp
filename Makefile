DB_URL=postgresql://root:secret@localhost:5433/fiber-workout?sslmode=disable

network:
	docker network create fitness-network

postgres:
	docker run --name postgres --network fitness-network -p 5433:5433 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.4-alpine3.18

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

createdb:
	docker exec -it postgres createdb --username=root --owner=root workout

dropdb:
	docker exec -it postgres dropdb workout

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateuplocal:
	migrate -path db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/fiber-workout?sslmode=disable" -verbose up

migratedownlocal:
	migrate -path db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/fiber-workout?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover -short ./...

server:
	go run main.go

dbdocs:
	dbdocs build loc doc/db.dbml

dbschema:
	dbml2sql --postgres - doc/schema.sql doc/db.dbml

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/DMonkey83/MyFitnessApp/db/sqlc Store

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=fitness_app \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: migrateuplocal migrateup migrateup1 migratedown sqlc test dbdocs dbschema proto evans
