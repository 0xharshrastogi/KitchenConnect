DSN:="sqlserver://sa:yourStrong(!)Password@localhost?database=KitchenConnectDB&sslmode=disable"
MIGRATION_DIR:=api/db/migrations

run-database:
	docker compose up database

run-api:
	cd api \
	&& go run .

stop:
	docker compose down -rm

migration-up:
	migrate -path $(MIGRATION_DIR) -database $(DSN) -verbose up

migration-down:
	migrate -path $(MIGRATION_DIR) -database $(DSN) -verbose down 1

migration-fix:
	migrate -path $(MIGRATION_DIR) -database $(DSN) -verbose force 0

migration-version:
	@migrate -path $(MIGRATION_DIR) -database $(DSN) version

create-migration:
	migrate create -ext sql -dir database/migration/ -seq init_mg
