# Path variable
MIGRATIONS_PATH = ./src/migrations
OPENAPI_PATH = ./docs/openapi.yaml
SWAGGER_COMPILED_PATH = ./docs/compile/swagger.yaml

serve:
	go run ./cmd/main.go

migration-generate:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

migration-up:
	go run ./cmd/migrate/migrate.go up

migration-down:
	go run ./cmd/migrate/migrate.go down

install-swagger-generate:
	npm install -g swagger-cli

swagger-generate:
	swagger-cli bundle $(OPENAPI_PATH) --outfile $(SWAGGER_COMPILED_PATH) --type yaml

gen-mock:
	for f in src/repo/*.go; do \
		base=$$(basename $$f .go | sed 's/\.repo//'); \
		mockgen -source=$$f -destination=src/repo/mocks/mock_$$base.go -package=mocks; \
	done

test:
	go test ./...

lint:
	golangci-lint run --fix