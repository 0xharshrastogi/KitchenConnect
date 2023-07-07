DSN:="sqlserver://sa:yourStrong(!)Password@localhost?database=KitchenConnectDB&sslmode=disable"
MIGRATION_DIR:="api/db/migrations"

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

migration-create:
	migrate create -ext sql -dir $(MIGRATION_DIR) remove_password_salt_column_from_users
