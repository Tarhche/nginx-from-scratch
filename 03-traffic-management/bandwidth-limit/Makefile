BUILD_NAME=bandwidth-limit
CONTAINER_NAME=bandwidth-limit

default: run

build:
	docker build -t ${BUILD_NAME} -f Dockerfile .

up start: build
	docker run --detach --rm --name=${CONTAINER_NAME} --publish 8050-8051:8080-8081 ${BUILD_NAME}

down stop:
	docker stop ${CONTAINER_NAME}

logs log:
	docker logs ${CONTAINER_NAME}

terminal sh:
	docker exec -it ${CONTAINER_NAME} sh

.PHONY: build up start down stop logs log terminal sh
