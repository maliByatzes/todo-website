tododb:
	psql -d tododb -U malib

test:
	go test -v -cover ./...

migrateup:
	migrate -path db/migration/ -database "postgresql://malib:Maliborh521908!@localhost:5432/tododb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://malib:Maliborh521908!@localhost:5432/tododb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: tododb test migrateup migratedown sqlc