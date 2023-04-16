compose-up:
	@docker-compose -f ./deployments/local/docker-compose.yaml --env-file ./deployments/local/.env up -d --build

compose-down:
	@docker-compose -f ./deployments/local/docker-compose.yaml --env-file ./deployments/local/.env down

compose-log-%:
	@docker-compose -f ./deployments/local/docker-compose.yaml logs -t --tail=100 -f $*

mongo-db:
	@docker run --name mongo-db -p 27017:27017 -v mongo-db:/data/db -d mongo:6.0
