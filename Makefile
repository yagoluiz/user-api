compose-build:
	@docker-compose -f ./deployments/local/docker-compose.yaml build --no-cache

compose-up:
	@docker-compose -f ./deployments/local/docker-compose.yaml --env-file ./deployments/local/.env up -d

compose-down:
	@docker-compose -f ./deployments/local/docker-compose.yaml --env-file ./deployments/local/.env down

compose-log-%:
	@docker-compose -f ./deployments/local/docker-compose.yaml logs -t --tail=100 -f $*