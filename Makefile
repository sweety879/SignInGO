migrateup:
migrate -path db/migrate -database "postgresql://sweety.seela:postgres@localhost:5432/credentials?sslmode=disable" -verbose up
migratedown:
migrate -path db/migrate -database "postgresql://sweety.seela:postgres@localhost:5432/credentials?sslmode=disable" -verbose down