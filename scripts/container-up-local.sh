#!/bin/sh
docker network create --driver bridge deploy_infra
docker network create --driver bridge infra
docker compose -f ./deploy/docker-compose.infra.local.yml up -d