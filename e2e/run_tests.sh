#!/bin/sh

# setup e2e testing env
docker compose up -d --build
docker compose exec -T test-db sh -c 'while ! mysqladmin ping -h0.0.0.0 --silent; do sleep 1; done'

# run e2e tests
go clean -testcache && go test ./e2e/tests/... -v

# print logs
echo "\n[CONTAINER LOGS]"
docker compose logs webhook-spaceway-lora

# taredown e2e testing env
docker compose down -v --rmi local

# cleanup
docker image prune -f
