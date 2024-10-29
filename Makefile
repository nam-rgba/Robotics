postgres:
	docker run --name fblv -e POSTGRES_USER=root -p 5431:5432 -e POSTGRES_PASSWORD=Namkhongbiet1 -d postgres

createdb:
	docker exec -it fblv createdb --username=root --owner=root  league

dropdb:
	docker exec -it fblv dropdb league

migrateup:
	migrate -path db/migration -database "postgres://root:Namkhongbiet1@localhost:5431/league?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:Namkhongbiet1@localhost:5431/league?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server