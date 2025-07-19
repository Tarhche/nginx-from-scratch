BUILD_NAME=response-compression
CONTAINER_NAME=response-compression

default: run

build:
	docker build -t ${BUILD_NAME} -f Dockerfile .

up start: build
	docker run --detach --rm --name=${CONTAINER_NAME} --publish 9055:9055 ${BUILD_NAME}

down stop:
	docker stop ${CONTAINER_NAME}

logs log:
	docker logs ${CONTAINER_NAME}

terminal sh:
	docker exec -it ${CONTAINER_NAME} sh

.PHONY: build up start down stop logs log terminal sh
