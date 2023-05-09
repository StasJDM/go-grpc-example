.SILENT:
include .env

mac-tools-install:
	brew install golang-migrate protobuf

migrate-create:
	migrate create -ext sql -dir db/migrations $(NAME)

migrate-up:
	migrate -database ${POSTGRES_URL} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRES_URL} -path db/migrations down