postgresdb:
	docker run --name postgresdb -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=p@ssw0rd -d postgres:16.2
createdb:
	 docker exec -it postgresdb createdb --username=root --owner=root moviedb
dropdb: 
	docker exec -it postgresdb dropdb moviedb
migrateup:
	migrate -path "./internal/adapter/db/sqlc/migrations" -database "postgresql://root:p@ssw0rd@localhost:5431/moviedb?sslmode=disable" -verbose up
migratedown:
	migrate -path "./internal/adapter/db/sqlc/migrations" -database "postgresql://root:p@ssw0rd@localhost:5431/moviedb?sslmode=disable" -verbose down 

	 

.PHONY: postgres createdb dropdb migrateup migratedown 