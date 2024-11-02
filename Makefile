
migrate_down:
	migrate -path ./migrations -database 'postgres://${PGSQL_USERS_USER}:${PGSQL_USERS_PASSWORD}@${PGSQL_USERS_HOST}:${PGSQL_USERS_PORT}/${PGSQL_USERS_DB_NAME}?sslmode=disable' down

migrate_up:
	migrate -path ./migrations -database 'postgres://${PGSQL_USERS_USER}:${PGSQL_USERS_PASSWORD}@${PGSQL_USERS_HOST}:${PGSQL_USERS_PORT}/${PGSQL_USERS_DB_NAME}?sslmode=disable' up


run_dev_app:
	go run main.go

run_app:
	./main

run: migrate_up run_app
