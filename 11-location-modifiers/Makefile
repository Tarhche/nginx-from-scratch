BUILD_NAME=location-modifiers
CONTAINER_NAME=location-modifiers

default: run

build:
	docker build -t ${BUILD_NAME} -f Dockerfile .

up start: build
	docker run --detach --rm --name=${CONTAINER_NAME} --publish 9091-9095:9091-9095 ${BUILD_NAME}

down stop:
	docker stop ${CONTAINER_NAME}

logs log:
	docker logs ${CONTAINER_NAME}

terminal sh:
	docker exec -it ${CONTAINER_NAME} sh

.PHONY: build up start down stop logs log terminal sh
