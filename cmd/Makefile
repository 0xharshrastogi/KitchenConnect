run-database:
	docker compose up database

run-api:
	cd api \
	&& go run .

stop:
	docker compose down -rm

migration-up:
	migrate -path ./api/db/migrations -database "sqlserver://sa:yourStrong(!)Password@localhost/KitchenConnectDB?sslmode=disable" -verbose up

create-migration:
	migrate create -ext sql -dir database/migration/ -seq init_mg
