BUILD_NAME=bandwidth-limit
CONTAINER_NAME=bandwidth-limit

default: run

up start:
	docker compose up -d --build

down stop:
	docker compose down --remove-orphans --volumes

logs log:
	docker compose logs proxy

terminal sh:
	docker compose exec proxy sh

.PHONY: up start down stop logs log terminal sh
