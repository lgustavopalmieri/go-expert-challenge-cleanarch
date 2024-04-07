include .env
export

createmigration:
ifeq ($(name),)
	@echo "Please, provide a migration name. Ex: make createmigration name=my_migration"
	@exit 1
endif
	migrate create -ext=sql -dir=internal/database/postgres/sql/migrations -seq $(name)

migrate:
	migrate -path=internal/database/postgres/sql/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5436/$(DB_NAME)?sslmode=disable" -verbose up
