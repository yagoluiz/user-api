compose-up:
	@docker-compose -f ./deployments/local/docker-compose.yaml --env-file ./deployments/local/.env up -d --build

compose-down:
	@docker-compose -f ./deployments/local/docker-compose.yaml --env-file ./deployments/local/.env down

compose-log-%:
	@docker-compose -f ./deployments/local/docker-compose.yaml logs -t --tail=100 -f $*

mongo-db:
	@docker run --name mongo-db -p 27017:27017 -v mongo-db:/data/db -d mongo:6.0

pkg-update:
	@go get -u all && go mod tidy

test-install-mocks:
	@sudo wget -P /tmp https://github.com/vektra/mockery/releases/download/v2.25.0/mockery_2.25.0_Linux_x86_64.tar.gz
	@sudo tar -xvzf /tmp/mockery_2.25.0_Linux_x86_64.tar.gz -C /tmp
	@sudo mv /tmp/mockery /usr/bin

test-generate-mocks:
	@mockery --srcpkg=./pkg/logger --name=Logger --filename=mocks/logger_mock.go --output ./pkg/ --outpkg mocks
	@mockery --srcpkg=./internal/repositories --name=UserRepositoryInterface --filename=mocks/user_repository_mock.go --output ./pkg/ --outpkg mocks
	@mockery --srcpkg=./internal/usecase --name=UserSearchUseCaseInterface --filename=mocks/user_usecase_mock.go --output ./pkg/ --outpkg mocks

test-run:
	@go test -coverpkg=./... -coverprofile=coverage-unit.out ./...

test-coverage:
	@go tool cover -html=coverage-unit.out
