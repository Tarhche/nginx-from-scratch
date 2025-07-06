BUILD_NAME=ab-testing-servers
CONTAINER_NAME=ab-testing-servers

default: run

build:
	docker build -t ${BUILD_NAME} -f Dockerfile .

up start: build
	docker run --detach --rm --name=${CONTAINER_NAME} --publish 8040:8000 ${BUILD_NAME}

down stop:
	docker stop ${CONTAINER_NAME}

logs log:
	docker logs ${CONTAINER_NAME}

terminal sh:
	docker exec -it ${CONTAINER_NAME} sh

.PHONY: build up start down stop logs log terminal sh
