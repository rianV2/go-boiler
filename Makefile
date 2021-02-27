APP_NAME=go-boiler
VERSION_VAR=main.Version
VERSION=$(shell git describe --tags)

build: dep
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X ${VERSION_VAR}=${VERSION}" -a -installsuffix nocgo -o ./bin ./...

dep:
	@echo ">> Downloading Dependencies"
	@go mod download

docker:
	@echo ">> Building Docker Image"
	@docker build -t ${APP_NAME}:latest .

run-api: dep
	@echo ">> Running API Server"
	@env $$(cat .env | xargs) go run github.com/remnv/go-boiler/cmd/go-boiler server

migrate: dep
	@echo ">> Running DB migration"
	@env $$(cat .env | xargs) go run github.com/remnv/go-boiler/cmd/go-boiler migrate

test-all: test-infra-up test-integration test-infra-down

test-integration: dep
	@echo ">> Running Integration Test"
	@env $$(cat .env.testing | xargs) env DB_MIGRATION_PATH=$$(pwd)/database/migration go test -tags=integration -count=1 -failfast -cover -covermode=atomic ./...

test-infra-up:
	$(MAKE) test-infra-down
	@echo ">> Starting Test DB"
	docker run -d --rm --name goboiler-test-mysql -p 3343:3306 --env-file .env.testing mysql:5.7
	docker cp $$(pwd)/deployments/docker goboiler-test-mysql:/tools
	docker exec goboiler-test-mysql sh -c '/tools/wait-for-mysql.sh 40'

test-infra-down:
	@echo ">> Shutting Down Test DB"
	@-docker kill goboiler-test-mysql

.PHONY: dep run-api migrate test-all test-integration test-infra-up test-infra-down 