migrate-create:
	migrate create -ext sql -dir migrations $(name)

migrate-up:
	migrate -path migrations -database "postgres://localhost:5432/rest_api?user=postgres&password=postgres2" up

migrate-down:
	migrate -path migrations -database "postgres://localhost:5432/rest_api?user=postgres&password=postgres2" down